# database
数据库
==

<pre>
data := make(map[string]string)
data[database.ProfileId] = "test"
data[database.ProfileDriver] = "mysql"
data[database.ProfileHost] = "127.0.0.1"
data[database.ProfileDatabase] = "test"
data[database.ProfileUsername] = "root"
data[database.ProfilePassword] = "123456"
data[database.ProfileWrite] = "true"

profile, err := database.NewProfile(data)
if err != nil {
    fmt.Println(err)
    return
}

builder := database.DriversBuilder{}

err2 := builder.AddProfile(profile)
if err2 != nil {
    fmt.Println(err2)
    return
}

drivers, err3 := builder.Build()
if err3 != nil {
    fmt.Println(err3)
    return
}

driver, err4 := drivers.GetWriter()
if err4 != nil {
    fmt.Println(err4)
    return
}
</pre>

<pre>
query := `DROP TABLE IF EXISTS user;`
_, err11 := database.Exec(driver, query)
if err11 != nil {
	fmt.Println(err11)
	return
}

query = `CREATE TABLE user (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
    phone varchar(32) NOT NULL DEFAULT '' COMMENT '手机号',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

_, err12 := database.Exec(driver, query)
if err12 != nil {
	fmt.Println(err12)
	return
}
</pre>

<pre>
r21, err21 := database.Exec(driver, "INSERT INTO user (name, phone) VALUES (?, ?);", "张三", "13000000001")
if err21 != nil {
	fmt.Println(err21)
	return
}

fmt.Println(r21)

r22, err22 := database.Exec(driver, "INSERT INTO user (name, phone) VALUES (?, ?);", "李四", "13000000002")
if err22 != nil {
	fmt.Println(err22)
	return
}

fmt.Println(r22)
</pre>

<pre>
r31, err31 := database.Exec(driver, "UPDATE user SET phone = ? WHERE id = ?", "18000000001", 1)
if err31 != nil {
    fmt.Println(err31)
    return
}

fmt.Println(r31)
</pre>

<pre>
r41, err41 := database.Exec(driver, "DELETE FROM user WHERE id = ?;", 2)
if err41 != nil {
    fmt.Println(err41)
    return
}

fmt.Println(r41)
</pre>

<pre>
r51, err51, closeErr51 := database.First(driver, "SELECT * FROM user WHERE id = ?", 1)
if err51 != nil {
    fmt.Println(err51)
    return
}

if closeErr51 != nil {
    fmt.Println(closeErr51)
    return
}

fmt.Println(r51)

r52, err52, closeErr52 := database.Find(driver, "SELECT * FROM user")
if err52 != nil {
    fmt.Println(err52)
    return
}

if closeErr52 != nil {
    fmt.Println(closeErr52)
    return
}

fmt.Println(r52)

r53, err53, closeErr53 := database.AggregateInt(driver, "SELECT COUNT(1) AS 'aggregate' FROM user")
if err53 != nil {
    fmt.Println(err53)
    return
}

if closeErr53 != nil {
    fmt.Println(closeErr53)
    return
}

fmt.Println(r53)
</pre>

<pre>
// 进程正常关闭前
err61 := drivers.Close()
fmt.Println(err61)
</pre>
