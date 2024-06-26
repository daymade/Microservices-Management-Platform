> 在实际的项目开发中，我们需要就产品细节跟产品经理、设计师、业务运营人员来回沟通，确定产品文档中未能在第一次描述时全部确定的细节，
> 
> 由于项目特殊，我这里简单的假设了一些业务背景，这只是为了减少和面试官中间的沟通损耗。

- 业务背景：

	- 业务: 我们假设每个 Service 都是一个后端 API 项目，包含了一系列 API 集合
	- 版本管理：Service 有版本管理，版本管理的粒度在 Service 级别而不是 API 级别，
      - 比如 `/v1` 的 Service 可能包含 10 个 API，`/v2` 的Service 可能包含 12 个 API。
      - 版本号可以是任何符合语义化版本的值，例如  `v1` 、`v2`，或者这种 `v2024-06-26` 也可以。
	- 多租户：只设计最核心的 Service Cards 服务，不需要进行跨区域和多租户设计，比如 Region、Tenant。
	- 权限控制：用户能看到自己的项目，**也可以**看到其他人的项目，实现用户维度的项目过滤不在这一期的考虑范围内。

- 功能需求：
	- 搜索：用户可以通过名字和描述搜索指定 Service，
	- 分页： 不需要能跳转到指定页数，只支持上一页和下一页
	- 开发者体验：
		- UI 相关: 需要支持 url 规则化，能通过 Uri 跳转到任何中间页面，例如：
			- `services` 是列表页面，如果输入了过滤条件则是 `services?query=name` 。
			- 通过 `services/contact-us` 或者  `services/locate-us` 可以直接跳转到某个 Service 的详情页面。
            - url 里不要有 id，使用 Service Name 作为 path。

- 非功能需求：

	- API 规范：我们设计符合 [Google API 规范](https://google.aip.dev/) 的 API。
    - 监控：使用 victoriametrics 和 grafana 进行监控。
    - 配置文件：使用静态配置文件，现阶段不实现动态配置中心

define the domain model of an api management platform,
- include the concept of [user,service,version,api]
- each service can be created by only one user
- each service has multiple version
- each service contains multiple apis, related with specific version

- 数据量：

	- 总 Service 数量：10 ～ 10000
	- 总用户数量：1000 以下
    - 每个用户能够创建的 Service 数量有限，最多创建 10 个 service。
    - 每个 Service 的版本数量：最多 10 个版本。

- 技术选型：

	- 搜索：由于现阶段数据量很小，我们不引入搜索引擎，直接在数据库上实现过滤。


