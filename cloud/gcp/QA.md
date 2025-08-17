# Google Cloud CLI (gcloud) Q&A

### Q: gcloud cmdline如何列出所有compute/region

A:
可以使用以下 `gcloud` 命令来列出所有可用的 Compute Engine 区域：

```sh
gcloud compute regions list
```

### Q: 东京输入哪个region

A:
东京对应的 region 是 `asia-northeast1`。

要从列表中精确筛选，需要对 `name` 字段进行过滤，而不是 `description`。正确的命令是：

```sh
gcloud compute regions list --filter="name=asia-northeast1"
```
*(之前尝试使用 `description~tokyo` 进行过滤失败，因为默认列表不包含 description 字段)*
