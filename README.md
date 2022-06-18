### 処理利流れ
1. cmd/api/main.go
2. handler
3. usecase
4. domain
5. infra

今回はusecase,handlerにinterface抽象化をしない
→handlerが2階層下のdomainに依存したりします
→DIもしない

### フォルダ構成

```
- handler
    ∟rest
      ∟[各種domain].go → 全てのCRUD関数記載(interface,structは一つ)
- usecase
    ∟[各種domain].go → 全てのCRUD関数記載(interface,structは一つ) ここでdb接続&閉じる
- domain
    ∟model
      ∟[各種domain].go
    ∟repository 
      ∟[各種domain]_repository.go → CRUD一つのinterfaceにまとめる([各種domain]Repository interface)
- infra
    ∟persistence
      ∟[各種domain].go → 全てのCRUD関数記載
```

### handler
- HTTPリクエストを受け取り、UseCase を使って処理を行い、結果を返す
- 外部にあるものがなんであれ、その差異を吸収して、ユースケースに伝えるのが役目
- HTTP通信以外でも対応できるように、本プロジェクトでは /handler/rest というふうにディレクトリを切っている(restAPIを明示)

### usecase
- システムのユースケースを満たす処理の流れを実装
### domain
- model
    - model定義

- reoisitory
    - UserRepository : User における Repository のインターフェース
    - -> 依存性逆転の法則により infra 層は domain層に依存
    - 呼び出すdomainロジック(interfaceを実装)

### infra
- domainロジックの技術的関心を記載
- package名をrepository という名前にしたいが domain 配下の repository とパッケージ名が被ってしまうため persistence で代替

