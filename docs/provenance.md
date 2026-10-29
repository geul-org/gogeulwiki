# 출처 기록

wikigeul이 만드는 문서마다 남기는 출처(문서 전체의 출처, 변환 도구)와 위키데이터 statement 출처(references)의 변환 규칙이다.

## 문서 전체의 출처 (HEAD)

HEAD에 문서 전체의 근거 Edge SOURCE(Target `#0`) 둘을 둔다.

1. **개체 판:** 원천 = LOCAL. BODY 앞머리에서 `TRIPLE P577 #r POSIX-TIMESTAMP-SECONDS(entity.modified)`(판 저장 시각)과 `TRIPLE P854 #r <https://www.wikidata.org/w/index.php?oldid=<entity.lastrevid>>`(판 영구 주소)로 기술한다. `modified`가 없으면 URL만이라 REFERENCE-NODE다. `lastrevid`가 없으면 이 SOURCE를 두지 않는다.
2. **덤프 파일:** 원천 = `REFERENCE-NODE #d #f <dump.url>`로 정의한 LOCAL. 그 SOURCE에 `QUALIFIER P4092`(파일 SHA-256, `dump.sha256`, SHA-256-DIGEST; 모르면 생략)와 `QUALIFIER P813`(받은 시각 `dump.downloaded_at`, POSIX-TIMESTAMP-SECONDS).

- 판과 덤프를 SOURCE 둘로 나눈 근거: 판은 내용의 원천이고(영구 주소로 같은 내용을 다시 받을 수 있다) 덤프는 그 판을 받은 파일이다. 파일 해시와 받은 시각은 덤프 파일에 관한 인용 쪽 성질이라 덤프 SOURCE의 QUALIFIER다. 한 SOURCE에 섞으면 해시가 판의 해시처럼 읽힌다. 덤프 SOURCE는 같은 실행의 모든 문서에서 같은 바이트다.
- 실행 옵션은 문서에 적지 않는다. 문서 내용을 바꾸는 옵션이 없다(`-schema`·`-dump-id`는 로컬 DB 위치이고 덤프는 SOURCE가 고정한다). 범위·DB·커밋은 실행 보고서 첫 줄에 남는다.

## 변환 도구 (TRANSFORMED-BY)

TRANSFORMED-BY 둘을 둔다. 도구 Reference는 geul.org 용어 정의 문서(EXTERNAL-ISSUER-1WORD, 도구 하나에 문서 하나)이고, 버전은 전체 커밋 ID다(SHA-1 저장소는 SHA1-DIGEST, SHA-256 저장소는 SHA-256-DIGEST; `port.CommitDigest`).

| 도구 | geul.org 용어 정의 문서 | 버전 |
|---|---|---|
| 변환 도구 wikigeul (이 저장소) | `0x0001000000000001` | `-commit` |
| GEUL 라이브러리 [gogeul](https://github.com/geul-org/gogeul) | `0x0001000000000002` | `-lib-commit` |

- 적재 도구(wikipg)는 적지 않는다. wikipg는 덤프를 무손실로 저장하므로 문서는 덤프·wikigeul·라이브러리로 정해진다. 적재 도구 커밋은 DB `dump.loader_commit`에 남는다.
- 커밋은 커밋된 소스여야 한다. 작업 트리에 커밋되지 않은 변경이 있으면 그 커밋 ID를 쓰지 말고 먼저 커밋한다(도구는 아직 검사하지 않는다).
- Go 워크스페이스로 다른 gogeul을 쓰면 `go.mod`가 고정한 버전과 다를 수 있다. 그래서 `-lib-commit`을 따로 받는다.

## statement 출처 (references)

statement의 출처는 근거 Edge SOURCE(`0x20`)로 옮긴다. 위키데이터 reference 하나가 SOURCE 하나다.

```
TRIPLE     #178   P19   Q42   Q350  normal
SOURCE     #1347  #178  Q5375741                     원천 기술 없이 P248 하나 → 원천 = 그 항목
QUALIFIER  #1348  P1417 #1347 "biography/Douglas-Adams"
TRIPLE     #1350  P50   #1349 Q18145749 normal       원천 기술이 있으면 원천 = LOCAL #1349
…
SOURCE     #1356  #178  #1349
QUALIFIER  #1357  P813  #1356 GREGORIAN-DATE(2013-12-07)
```

- 읽는 표: `statement_reference`(statement_id, ordinal → hash)와 `reference_snak`(hash, dump_id, ordinal). 두 표의 기본 키 색인을 탄다(hash를 함수로 감싸지 않는다). 출처 순서는 `ordinal`이다.
- 출처 항목 순서는 그 출현의 snaks-order다. `reference_snak_order`가 있으면 출현별 순서를, 없으면 hash의 첫 출현 순서를 쓴다([database.md](database.md#출현별-출처-항목-순서)).
- 항목은 qualifier와 같은 규칙으로 값을 바꾼다(somevalue·novalue는 자리 표지). 옮기지 못한 항목은 보고서 `skipped`에 `reference ` 접두로 센다.
- **속성 분류**(`port.originProperties`): 원천 쪽은 P248(stated in), P854(reference URL), P1476(title), P1680(subtitle), P50(author), P2093(author name string), P98(editor), P655(translator), P577(publication date), P1433(published in), P123(publisher), P291(place of publication), P407(language of work or name). 나머지(P813, P143, P887, P4656, P304, 외부 식별자 등)와 분류되지 않은 속성은 인용 쪽이다.
- **원천 결정**(`port.originOf`):
  1. 원천 쪽 항목이 없으면(P143·P813만 등) 원천은 ABSENT(적지 않음). 변환기는 원천을 적지 않았을 뿐 "원천이 있지만 모른다"(UNKNOWN-VALUE)를 단언하지 않는다.
  2. 원천 쪽 항목이 P248 하나이고 값이 항목이면 원천은 그 WIKIDATA-QID.
  3. 그 밖에는 LOCAL 노드. 원천 쪽 항목을 속성 번호 순(같은 속성 안은 덤프 순)으로 정렬해 TRIPLE(P, LOCAL, 값, normal)로 쓴다. P248도 `TRIPLE P248 #L Q…`로 남는다. 정렬한 항목이 P854 하나이고 URI면 TRIPLE 대신 REFERENCE-NODE(LOCAL → URI)로 정의한다.
- 문서 안에서 원천 쪽 항목 집합이 같은 LOCAL은 한 번만 선언하고 공유한다(처음 쓰는 SOURCE 바로 앞). LOCAL 번호는 다시 만들 때 같은 기술의 노드가 이어받는다(`delta.LocalIDs.Node`, URL 원천은 `Define`).
- 인용 쪽 항목은 SOURCE 뒤에 QUALIFIER로 그 출현의 덤프 순서대로 둔다.
- **역변환 검사:** SOURCE마다 원천(항목이면 P248=항목, LOCAL이면 그 TRIPLE과 REFERENCE-NODE(P854), ABSENT면 없음)과 SOURCE의 QUALIFIER를 모으면 출처의 항목 집합이 된다. 출처 순서와 속성 사이 항목 순서는 잃는다. 같은 속성 안 값 순서는 남는다(그 출현의 순서. 다만 문서 안에서 공유하는 원천 LOCAL은 처음 선언한 출현의 순서로 기술된다). `BuildDoc`이 문서마다 역변환해 원래 집합과 대조하고, 다른 TRIPLE 수를 보고서 `ref_mismatch`에 적는다(0이어야 한다).
- 보고서 열: `sources`(SOURCE 수), `origin_qid`·`origin_absent`·`origin_local`(원천 종류별 수, LOCAL은 문서 안 선언 수), `origin_edges`(LOCAL을 기술하는 TRIPLE·REFERENCE-NODE 수), `source_qualifiers`, `ref_mismatch`.
