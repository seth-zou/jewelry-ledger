# Contract Upgrade Notes

1. Build and test the pure domain package.
2. Review state schema compatibility.
3. Package chaincode in a release pipeline.
4. Approve organization lifecycle changes according to the endorsement policy.
5. Roll back by stopping writes and returning clients to the previous approved definition when possible.
