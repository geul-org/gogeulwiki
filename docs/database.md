# 읽는 DB 테이블

wikigeul은 [wikipg](https://github.com/geul-org/wikipg)가 적재한 스키마(기본 `wikidata`)를 읽기만 한다. 테이블 정의는 wikipg `sql/wikidata.sql`이다.

| 테이블 | wikigeul이 쓰는 열 |
|---|---|
| `dump` | `url`, `sha256`, `downloaded_at`, `loader`, `loader_commit` (출처 기록) |
| `entity` | `entity_id`(있는지), `lastrevid`·`modified`(개체 판, 출처 기록), `datatype`(속성) |
| `label`, `description`, `alias` | `language`, `value` (모든 언어, 언어 코드 바이트 순. 별칭은 그 안에서 `ordinal` 순) |
| `statement` | `statement_id`, `property`, `rank`, `snaktype`, `value_type`, `value`(jsonb), `value_entity` |
| `qualifier` | `statement_id`, `ordinal`(qualifiers-order), `property`, `snaktype`, `value_type`, `value`, `value_entity` |
| `statement_reference`, `reference_snak` | 출처: statement의 출처 순서·hash, 출처 항목(속성, 값 열은 qualifier와 같다) |
| `reference_snak_order` (있으면) | 출현별 출처 항목 순서: `(statement_id, ordinal)` → `perm`(wikipg `build-reference-order`). 표가 없거나 그 출현에 행이 없으면 `reference_snak.ordinal` 순서 |
| `sitelink` | `site`, `title`(`title_json`), `badges` (사이트 ID 바이트 순 = 덤프 키 순) |

## 읽기 규칙

- 모든 질의는 `dump_id`로 한 판만 읽는다(`-dump-id`).
- `statement_id`는 wikipg가 매긴다. 속성 사이는 속성 ID 문자열 순, 같은 속성 안은 덤프 순이다. 위키데이터에서 속성 사이 순서는 뜻이 없다.
- 값은 개체면 맨 ID(`value_entity`), 문자열이면 따옴표 없는 문자열, 그 밖은 datavalue JSON(jsonb의 공백·키 순서)이다.
- `\u0000`이 든 값은 wikipg가 `value_json`(레이블·별칭·설명도 `value_json`)에 원문 JSON으로 두므로 그것을 읽는다(`source.exactValue`).
- novalue·somevalue는 값이 비어 있고 `snaktype` 열로 구별한다(자리 표지로 옮긴다, [변환 규칙](conversion.md)). 값 snak인데 값이 빈 것만 건너뛴다(보고서 `빈 값`).
- 본 값과 qualifier 값은 같은 형태다.

## 적재 도구 확인

`dump.loader`가 wikipg(`github.com/geul-org/wikipg/cmd/wikipg`, 또는 wikipg의 옛 이름으로 기록된 경로)가 아니면 실행을 거부한다. wikipg는 덤프를 무손실로 저장하므로(덤프 줄 원문 `entity_raw`, 구조화 표가 못 담는 값은 `*_json` 열) 문서 내용은 덤프·wikigeul·GEUL 라이브러리로 정해진다. 다른 적재 도구는 무손실을 알 수 없다.

## 출현별 출처 항목 순서

- `reference_snak`은 hash마다 덤프상 첫 출현의 순서 하나만 둔다. 위키베이스 출처 hash는 snak 순서와 무관해서, 같은 hash의 다른 출현은 순서가 다를 수 있다.
- 그런 출현은 wikipg의 `reference_snak_order`가 순열 `perm`(reference_snak ordinal을 그 출현 순서로)을 둔다. wikigeul은 `LEFT JOIN LATERAL`로 PK를 한 번 조회해 `coalesce(array_position(perm, ordinal) - 1, ordinal)` 순으로 읽는다(`source.refOrderSQL`).
- 행이 없는 출현은 `reference_snak.ordinal` 순서다. 표가 없는 스키마(옛 적재)도 그렇다. `source.LoadDump`가 `to_regclass`로 표가 있는지 보고, wikigeul이 실행 첫 줄(stdout)에 `reference snak order: per occurrence (reference_snak_order) true|false`로 알린다.

## 세션

DB는 연결 하나에 `SET jit = off`로 쓴다. 출처 질의(statement_reference ⋈ reference_snak)의 행 수를 플래너가 크게 잡아 JIT 컴파일이 개체마다 수백 ms 걸리기 때문이다(끄면 1ms 안쪽).
