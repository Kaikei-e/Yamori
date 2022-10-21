# 宿泊管理サービス

Yamori Roster

<br>

## 概要

- 運営している宿に関しての情報を管理できる。
  - 宿の名前や住所、電話番号、メールアドレス、キャパシティなどの情報を管理できる。
- 予約の情報を確認することができる。
  - 最小で 1 時間単位で、宿泊客の管理を行うことができる。
  - 1 時間、1 日、1 週間、1 ヶ月、1 年の予約状況を、柔軟に確認することができる。
- 宿泊客の予約をキャンセルすることができる。
- 各宿には、組織の概念が存在し、権限によって行えることが異なる。
  - 宿のオーナーには、従業員の権限管理を含めた、その組織へ与えられた全権限が付与される。
  - また、マネージャー、従業員には閲覧権限と管理のみが割与えられ、アルバイト社員には閲覧権限のみが与えられる。
- 組織がサブスクリプションに加入している場合には、宿泊客や宿泊コースなどに関して、より詳細でまとまった統計情報を見ることができる。
- オーナーと、マネージャーは、休日などを設けて、予約を受け付けない日を決めることができる。

<br>

## 機能 詳細

<br>

### 宿の管理

- 宿の名前や住所、電話番号、メールアドレス、キャパシティなどの情報を編集、公開できる。
- グループ会社である場合には、グループ会社の情報を編集し、組織管理サービスへ申請を行うことで、グループ会社として登録することができる。
  - 組織管理サービスでの承認があった場合には、グループ会社として登録される。

<br>

### 予約の管理

- カレンダーにて、予約の状況を確認することができる。
  - 基本的に日数での管理ではなく、月単位での管理を行う。
    - 1月なら31日間、2月なら28日間、3月なら31日間、4月なら30日間のように、月単位での管理を行う。
- 予約の情報を確認することができる。
  - 最小で 1 時間単位で、宿泊客の管理を行うことができる。
  - 1 時間、1 日、1 週間、1 ヶ月、1 年の予約状況を、柔軟に確認することができる。
- 宿泊客の予約をキャンセルすることができる。
  - たとえば、悪天候などで宿泊客の安全が確保できないと言ったシナリオなど。
- 部屋の空き数を確認できる。

<br>

### 人員の管理

- 組織の概念が存在し、権限によって行えることが異なる。
  - 宿のオーナーには、従業員の権限管理を含めた、その組織へ与えられた全権限が付与される。
  - また、マネージャー、従業員には閲覧権限と管理のみが与えられ、アルバイト社員には閲覧権限のみが与えられる。
  - アカウントの発行、停止、復活、削除ができる。
  - 加えて、権限の追加、削除のコントロールができる。

<br>

### 統計情報

- サブスクに加入していなくても、基本的な宿泊客の推移、年齢、性別などの統計情報を見ることができる。
- 組織がサブスクリプションに加入している場合には、宿泊客や宿泊コースなどに関して、より詳細でまとまった統計情報を見ることができる。
  - たとえば、宿泊コースや日数、オプションの選択数など。

<br>

### 休日の設定

- オーナーと、マネージャーは、休日などを設けて、予約を受け付けない日を決めることができる。
  - たとえば、オーナーが、毎週水曜日を休日として設定した場合、その日は、予約を受け付けない。
  - また、オーナーが、毎年1月1日を定休日として設定した場合、その日は、予約を受け付けない。