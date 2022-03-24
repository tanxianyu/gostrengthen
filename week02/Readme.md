
- 不应该在dao去Wrap error，dao层查询数据库，返回结果和根error即可
- 可以在dao层的上一层去处理这个Wrap error，dao的上一层才会对外接入
