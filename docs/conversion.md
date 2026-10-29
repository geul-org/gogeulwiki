# 변환 규칙

위키데이터 개체 하나를 GEUL 문서 하나로 옮기는 규칙이다. 명세는 [geul.org 문법](https://geul.org/ko/grammar/)(SSOT)을 따른다.

## 문서

- **Issuer:** `wikidata.geul.org` = `0x00010000` 하나. 위키데이터의 모든 개체 종류를 한 Issuer 아래 둔다. Q·P·L은 같은 출처를 같은 주체가 옮긴 것이라 나누지 않는다.
- **문서 단위:** 개체(QID·PID) 하나당 `.geul` 문서 하나. 덤프에 있는 모든 언어(`mul` 포함)를 한 문서에 담는다.
- **문서 ID:** 64비트 중 맨 위 바이트가 개체 종류의 ASCII 글자, 아래 56비트가 번호다(`port.DocNumber`). hex만 보고 종류를 읽을 수 있다. 번호가 2^56 이상이면 오류다.

  | 종류 | 맨 위 바이트 | 예 |
  |---|---|---|
  | 항목 Q | `0x51` | Q42 → `0x510000000000002a` |
  | 속성 P | `0x50` | P31 → `0x500000000000001f` |
  | Lexeme L | `0x4C` | L42 → `0x4c0000000000002a` (예약. Lexeme 문서는 아직 만들지 않는다) |
  | Entity Schema E | `0x45` | E10 → `0x450000000000000a` (예약) |

- **이름 색인:** Issuer 디렉터리의 `documents.geul` 하나. 이름은 소문자 종류 글자 + 번호다(`q42`, `p31`). 색인은 이번 실행분과 그 Issuer 디렉터리에 이미 공표된 문서를 모두 담는다(범위를 나눠 실행해도 앞 실행의 이름이 지워지지 않는다).
- Q·P를 가리키는 Reference(WIKIDATA-QID·PID)는 문서 위치와 상관없다.
- **실행 순서:** Q 범위를 먼저 옮기고, 그 문서들이 쓴 속성(Predicate, 값이 속성 개체인 것)을 이어서 옮긴 뒤 `documents.geul`을 쓴다.
- **공표:** 지금은 `resolve.Dir.Overwrite`로 실행할 때마다 v1을 덮어쓰고 버전 목록도 v1 하나만 둔다. 내용이 같으면 아무 파일도 바꾸지 않는다.

## 레코드 순서

HEAD의 출처 SOURCE와 그 원천 기술(BODY 앞머리), LABEL, DESCRIPTION, ALIAS, (속성이면) 데이터 타입, TRIPLE과 그 QUALIFIER·SOURCE, sitelink 순이다. 출처 기록은 [provenance.md](provenance.md).

- 레이블·설명·별칭은 LABEL·DESCRIPTION·ALIAS(Edge Id, Subject, Object)이고 Object는 UTF8-LANGUAGE-TEXT, 언어는 LANGUAGE-BCP47-CODE다. 빈 설명은 건너뛴다. 언어 코드는 BCP 47 형식만 확인한다(아래 언어 코드 대응).
- 클레임은 TRIPLE(Edge Id, Predicate, Subject, Object, Rank)이고, 그 qualifier는 바로 뒤에 QUALIFIER(Edge Id, Property, 대상 TRIPLE의 Edge Id, Value)로 잇는다. qualifier 값은 본 값과 같은 규칙으로 바꾼다. 그 뒤에 출처마다 SOURCE와 그 QUALIFIER가 온다.
- qualifier 순서는 `qualifier.ordinal`(위키데이터 `qualifiers-order`, 원본 그대로)이다.

## Edge Id 유지

Edge Id(LOCAL ID)는 버전을 넘어 유지한다. 다시 생성할 때 직전 공표본과 그 버전 목록의 `NEXT-LOCAL-ID`를 읽어(`resolve.Dir.Current`), 내용이 같은 서술은 같은 번호를, 바뀌거나 새로 생긴 서술은 다음 번호를 받는다(gogeul `delta.LocalIDs`). 사라진 번호는 다시 쓰지 않는다. 내용이 바뀐 클레임은 "삭제 + 신규"가 된다. HEAD Edge도 BODY와 한 번호 공간을 쓰고, 번호는 HEAD가 먼저 받는다.

## 값 변환

| 위키데이터 | GEUL |
|---|---|
| 항목 `Q…`, 속성 `P…` | WIKIDATA-QID, WIKIDATA-PID Reference |
| Lexeme `L…`, EntitySchema `E…` | 건너뜀 (Reference Codec 없음) |
| string, external-id, math, musical-notation | UTF8-TEXT (external-id는 식별자 텍스트 그대로) |
| url | URI. 비ASCII·URI에 못 쓰는 바이트는 퍼센트 인코딩(RFC 3987 3.1, 되돌릴 수 있다). 스킴이 없으면 원문 UTF8-TEXT |
| commonsMedia | URI `https://commons.wikimedia.org/wiki/File:` + 파일 이름(공백 → `_`, 한 경로 조각으로 퍼센트 인코딩) |
| geo-shape, tabular-data | URI `https://commons.wikimedia.org/wiki/` + 페이지 이름(`Data:…`), 위와 같은 인코딩 |
| time, 그레고리력 일 정밀도(기원전 포함) | GREGORIAN-DATE |
| time, 그레고리력 월·연·10년 정밀도 | EDTF |
| time, 율리우스력 일·월·연 정밀도 | GEDCOM-DATE (`JULIAN 14 OCT 1066`) |
| time, 그 밖(세기 이상 정밀도, 시간대·오차 범위 등) | datavalue JSON 원문을 UTF8-TEXT로. 날짜는 건너뛰지 않는다 |
| somevalue (본 값, qualifier) | 자리 표지 UNKNOWN-VALUE `0x0001` (값이 있으나 무엇인지 모른다). rank·qualifier 순서는 그대로 |
| novalue (본 값, qualifier) | 자리 표지 NO-VALUE `0x0002` (그 qualifier 범위 안에서 값이 없다). 본 값 novalue의 값 qualifier(P580·P582 등)도 함께 옮긴다 |
| monolingualtext | UTF8-LANGUAGE-TEXT (언어는 LANGUAGE-BCP47-CODE, 아래 언어 코드 대응). 그래도 BCP 47 형식이 아니면(빈 언어) datavalue 원문을 UTF8-TEXT로 |
| globecoordinate | 지구(Q2)는 COORDINATE-EARTH, 다른 천체는 COORDINATE-GEO(천체 = WIKIDATA-QID). precision은 환산 없이 불확도(binary32)와 단위 QID로 옮긴다: 도 Q28390·각분 Q209426·각초 Q829073 순서로, 역변환(`port.precisionText`)이 precision 원문을 바이트 그대로 되돌리는 첫 단위(`port.coordinatePrecision`). precision 없음은 미상(단위 ABSENT). 원문을 되돌리지 못하는 precision·천체 없음·고도 있음·범위 밖은 datavalue 원문을 UTF8-TEXT로 |
| quantity | QUANTITY(amount·단위·하한·상한). 수와 경계는 자릿수 그대로, 단위 `1`은 무단위(ABSENT 표지), 위키데이터 항목 단위는 WIKIDATA-QID. 그 밖의 단위·형식은 원문 UTF8-TEXT |
| 그 밖(Lexeme 값 등) | 건너뜀. "옮기지 못함"은 표지로 쓰지 않는다(원천에는 값이 있으므로 UNKNOWN-VALUE가 아니다) |

## 언어 코드 대응

위키미디어 언어 코드 중 유효한 BCP 47이 아닌 14개는 BCP 47 tag로 옮긴다(`port.wikimediaLanguage`, 대소문자 무시). 나머지 코드는 그대로 쓴다. 근거는 IANA Language Subtag Registry 항목이다(File-Date 2026-09-17, IANA 프로토콜 레지스트리, CC0 1.0: [IANA licensing terms](https://www.iana.org/help/licensing-terms)).

| 위키미디어 | BCP 47 | IANA 근거 |
|---|---|---|
| `sr-el` | `sr-Latn-x-srel` | language `sr`, script `Latn` + 사설 확장 |
| `bat-smg` | `sgs-x-batsmg` | language `sgs`(Samogitian) + 사설 확장 |
| `fiu-vro` | `vro-x-fiuvro` | language `vro`(Võro) + 사설 확장 |
| `cbk-zam` | `cbk-x-cbkzam` | language `cbk`(Chavacano) + 사설 확장 |
| `simple` | `en-simple` | language `en`, variant `simple`(Simplified form) |
| `eml` | `roa-x-eml` | language `roa`(Romance languages, collection) + 사설 확장. `eml`은 레지스트리에 없다 |
| `roa-rup` | `rup-x-roarup` | language `rup`(Aromanian) + 사설 확장 |
| `zh-classical` | `lzh-x-zhclass` | language `lzh`(Literary Chinese) + 사설 확장 |
| `map-bms` | `jv-x-bms` | language `jv` + 사설 확장 |
| `roa-tara` | `nap-x-tara` | language `nap` + 사설 확장 |
| `de-formal` | `de-x-formal` | language `de` + 사설 확장 |
| `es-formal` | `es-x-formal` | language `es` + 사설 확장 |
| `hu-formal` | `hu-x-formal` | language `hu` + 사설 확장 |
| `nl-informal` | `nl-x-informal` | language `nl` + 사설 확장 |

- **일대일:** 서로 다른 위키미디어 코드는 서로 다른 tag가 된다. 위키데이터는 `sgs`와 `bat-smg`, `vro`와 `fiu-vro`, `rup`와 `roa-rup`, `lzh`와 `zh-classical`, `egl`와 `eml`을 모두 용어 언어로 받고, `sr-latn`와 `sr-el`, `cbk`와 `cbk-zam`을 모두 monolingualtext 언어로 받는다. 그래서 대상 tag가 그 자체로 위키데이터 코드인 경우, BCP 47이 아닌 쪽에 사설 확장을 붙인다. 사설 확장 하위 태그는 위키미디어 코드에서 하이픈을 뺀 것이다(8자를 넘는 `zh-classical`만 `zhclass`).
- **`eml`:** 에밀리아·로마냐어 전체를 뜻한다. ISO 639-3이 2009년에 `egl`(에밀리아어)·`rgn`(로마냐어)로 나누며 폐지했고, IANA 레지스트리에는 처음부터 없다. 에밀리아어 `egl`로 옮기면 뜻이 좁아지므로, 둘을 함께 덮는 등록된 가장 좁은 언어 하위 태그인 로망스어군 `roa`에 `x-eml`을 붙인다.
- 사설 확장(`x-`)의 뜻: `x-formal` 존댓말(격식체), `x-informal` 반말(비격식체), `x-bms` 자바어 반유마스 방언, `x-tara` 나폴리어 타란토 방언, `x-eml` 에밀리아·로마냐어. `x-srel`·`x-batsmg`·`x-fiuvro`·`x-cbkzam`·`x-roarup`·`x-zhclass`는 "그 위키미디어 코드로 적힌 텍스트"라는 뜻이다(언어는 `x` 앞 tag와 같다).
- 표에 없고 BCP 47 형식도 아닌 코드(빈 언어 등)는 레이블·설명·별칭이면 건너뛰고(보고서 `skipped`), monolingualtext면 원문 UTF8-TEXT로 보존한다(보고서 `raw`).

### 역대응: BCP 47 tag → 위키미디어 코드

역대응은 유일하다(`port.wikimediaCode`). tag가 위 표의 BCP 47 열에 있으면(대소문자 무시) 그 행의 위키미디어 코드이고, 없으면 tag 자체가 위키미디어 코드다. 위 표는 위키데이터가 받는 모든 코드(용어 언어와 monolingualtext 언어 합집합 1,488개, 2026-10-29 `wbcontentlanguages`) 위에서 일대일이고, 표의 BCP 47 열 tag는 어느 것도 위키데이터 코드가 아니다. 그래서 어떤 코드든 tag로 옮겼다가 되돌리면 원래 코드가 나온다. 이 목록은 `testdata/wikidata-languages.txt`에 두고 `TestWikimediaLanguageInjective`가 검사한다. 위키데이터가 새 코드를 받으면 목록을 갱신하고 다시 검사한다.

| BCP 47 | 위키미디어 | | BCP 47 | 위키미디어 |
|---|---|---|---|---|
| `sr-Latn-x-srel` | `sr-el` | | `lzh-x-zhclass` | `zh-classical` |
| `sgs-x-batsmg` | `bat-smg` | | `jv-x-bms` | `map-bms` |
| `vro-x-fiuvro` | `fiu-vro` | | `nap-x-tara` | `roa-tara` |
| `cbk-x-cbkzam` | `cbk-zam` | | `de-x-formal` | `de-formal` |
| `en-simple` | `simple` | | `es-x-formal` | `es-formal` |
| `roa-x-eml` | `eml` | | `hu-x-formal` | `hu-formal` |
| `rup-x-roarup` | `roa-rup` | | `nl-x-informal` | `nl-informal` |
| 그 밖의 tag | tag 그대로 | | | |

## 속성 데이터 타입

속성(P) 문서는 그 속성의 데이터 타입(`entity.datatype`)을 문서 안에 적는다. 다른 문서에서 그 속성을 만난 해석기가 속성 문서 하나만 보면 되게 한다.

```
REFERENCE-NODE( #n → URI http://wikiba.se/ontology#propertyType )
TRIPLE( #n, P214, URI http://wikiba.se/ontology#ExternalId, normal )
```

어휘는 위키데이터 RDF의 Wikibase 온톨로지 이름 그대로다(18종, `port.wikibaseType`). 모르는 데이터 타입은 쓰지 않고 보고서 `skipped`에 남긴다.

## Sitelink

위키데이터 RDF가 sitelink를 쓰는 모양(문서 페이지가 주어, `schema:about` 개체, `wikibase:badge` 배지) 그대로 옮긴다. 쓰는 Edge는 REFERENCE-NODE와 TRIPLE뿐이다.

```
REFERENCE-NODE( #a → URI http://schema.org/about )              문서마다 한 번, 첫 페이지 앞
REFERENCE-NODE( #p → URI https://en.wikipedia.org/wiki/Douglas_Adams )
TRIPLE( #a, #p, Q42, normal )                                    페이지 #p는 Q42에 관한 것이다
REFERENCE-NODE( #b → URI http://wikiba.se/ontology#badge )       배지가 처음 나올 때 한 번
TRIPLE( #b, #p, Q17437796, normal )                              #p의 배지(알찬 글)
```

- REFERENCE-NODE(Q42 → 페이지 URI)로 쓰지 않는다. REFERENCE-NODE는 Subject가 Object를 지시한다는 바인딩이라, Q42(사람)가 위키백과 문서를 지시한다는 뜻이 되고 한 Reference가 사이트 수만큼 다른 Object에 묶인다.
- 순서: 사이트 ID 바이트 순(덤프 JSON의 키 순), 배지는 덤프 순. 사이트 ID·언어·이름은 URL에서 되살릴 수 있어 따로 적지 않는다.
- URL: `https://` + 호스트 + `/wiki/` + 제목. 제목은 공백을 `_`로 바꾸고 MediaWiki `wfUrlencode`처럼 인코딩한다(영숫자와 `-._;@$!*(),/~:`는 그대로, 나머지 바이트는 대문자 `%XX`).
- 호스트: 사이트 ID = 언어 부분 + 프로젝트 접미사. 접미사 `wiki`→`wikipedia.org`, `wiktionary`·`wikibooks`·`wikinews`·`wikiquote`·`wikisource`·`wikiversity`·`wikivoyage`→같은 이름 `.org`, 언어 부분의 `_`는 `-`(`zh_min_nanwiki` → `zh-min-nan.wikipedia.org`). 규칙과 다른 사이트는 표로 둔다(Wikimedia SiteMatrix와 대조, `port.wikiSiteHost`):

  | 사이트 ID | 호스트 |
  |---|---|
  | `be_x_oldwiki` | `be-tarask.wikipedia.org` |
  | `commonswiki` | `commons.wikimedia.org` |
  | `foundationwiki` | `foundation.wikimedia.org` |
  | `mediawikiwiki` | `www.mediawiki.org` |
  | `metawiki` | `meta.wikimedia.org` |
  | `outreachwiki` | `outreach.wikimedia.org` |
  | `sourceswiki` | `wikisource.org` |
  | `specieswiki` | `species.wikimedia.org` |
  | `wikidatawiki` | `www.wikidata.org` |
  | `wikifunctionswiki` | `www.wikifunctions.org` |
  | `wikimaniawiki` | `wikimania.wikimedia.org` |

- 규칙에도 표에도 맞지 않는 사이트는 건너뛰고 보고서 `skipped`에 `sitelink 사이트: <ID>`로 센다. 항목 ID가 아닌 배지는 `sitelink badge: <값>`.
