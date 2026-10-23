# jewelry-ledger

## 项目定位

Go-based Fabric chaincode and ledger governance repository for jewelry trace evidence.

## 标准命令

- `go test ./...`
- `go vet ./...`

## 约束

- 不提交生产 MSP、私钥、证书或真实组织凭据。
- Chaincode 逻辑应先在纯领域层测试，再接 Fabric 适配层。
- 网络目录只存模板、策略和治理说明。
