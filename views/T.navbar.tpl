{{define "navbar"}}
<nav class="navbar navbar-light navbar-expand-lg" style="background-color: #e3f2fd;" >
    <div class="container-fluid">
        <a class="navbar-brand" href="/"><img src="/static/img/code.png" width="35" height="35" >        JRPDShang</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#Content" aria-controls="Content" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="Content">
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                {{if .IsHome}}
                <li  class="nav-item" ><a  class="nav-link active" href="/"><img src="/static/img/home.png" width="20" height="20" class="d-inline-block align-text-top">  首页</a></li>
                {{else}}
                <li  class="nav-item" ><a  class="nav-link" href="/"><img src="/static/img/home.png" width="20" height="20" class="d-inline-block align-text-top">  首页</a></li>
                {{end}}
                {{if .IsCategory}}
                <li class="nav-item"><a class="nav-link active" href="/category"><img src="/static/img/category.png" width="20" height="20" class="d-inline-block align-text-top">  分类</a></li>
                {{else}}
                <li class="nav-item"><a class="nav-link" href="/category"><img src="/static/img/category.png" width="20" height="20" class="d-inline-block align-text-top">  分类</a></li>
                {{end}}
                {{if .IsTopic}}
                <li class="nav-item"><a class="nav-link active" href="/topic"><img src="/static/img/topic.png" width="20" height="20" class="d-inline-block align-text-top">  文章</a></li>
                {{else}}
                <li class="nav-item"><a class="nav-link" href="/topic"><img src="/static/img/topic.png" width="20" height="20" class="d-inline-block align-text-top">  文章</a></li>
                {{end}}
                {{if .IsUser}}
                <li class="nav-item"><a class="nav-link active" href="/user"><img src="/static/img/user.png" width="20" height="20" class="d-inline-block align-text-top">  用户</a></li>
                {{else}}
                <li class="nav-item"><a class="nav-link" href="/user"><img src="/static/img/user.png" width="20" height="20" class="d-inline-block align-text-top">  用户</a></li>
                {{end}}
            </ul>
            <form class="d-flex" action="/" method="get">
                <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search" name="label" value="{{.Label}}">
                <button class="btn btn-outline-success" type="submit">Search</button>
            </form>
{{end}}