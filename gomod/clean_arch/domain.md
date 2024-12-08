# Доменная область

## [<<< ---](../clean_arch.md)

Домен ничего не знает о других слоях. Он содержит чистую бизнес-логику.

Entities содержат бизнес-правила, независимые от приложения. И они *не просто объекты с данными. Entities могут содержать ссылки на объекты с данными, но основное их назначение в том, чтобы реализовать методы бизнес-логики, которые могут использоваться в различных приложениях*.