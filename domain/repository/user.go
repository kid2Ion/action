package repository

// UserRepository : User における Repository のインターフェース
//  -> 依存性逆転の法則により infra 層は domain 層（本インターフェース）に依存
// 呼び出すdomainロジック(interfaceを実装)
