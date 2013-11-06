{{define "navbar"}}
<a class="navbar-brand" href="/">锅巴的博客</a>
<div>
    <ul class="nav navbar-nav">
        <li {{if .IsHome}} class="active"{{end}}><a href="/">首页</a></li>
        <li {{if .IsCategory}} class="active"{{end}}><a href="/category">分类</a></li>
        <li {{if .Topic}} class="active"{{end}}><a href="/topic">文章</a></li>
    </ul>
</div>

<div class="pull-right">
	<wl class="nav navbar-nav">
		{{if .IsLogin}}
		<li {{if .IsLoginPage}} class="active" {{end}}><a href="/login?exit=true">退出</a></li>
		{{else}}
		<li {{if .IsLoginPage}} class="active" {{end}}><a href="/login">登录</a></li>
		{{end}}
	</wl>
</div>
{{end}}