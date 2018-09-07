<html>
<head>
    <title>user info</title>
</head>
<body>
    <div>
        <form action="user" method="post">
            all: {{.}}<br>
            hello, {{if gt .age 18}}man {{else}} boy {{end}} {{.name}}! <br>
            username: <input type="text" name="username" value="{{.name}}"><br>
            age: <input type="text" name="age" value="{{.age}}"><br>
            sex: <input type="text" name="sex" value="{{.sex}}"><br>
            <button>submit</button>
        </form>
    </div>
</body>
</html>