# 导入 requests 包
import requests
import re
import pymysql
import json

# 发送请求
res = requests.get('https://www.jinjing365.com/index.asp')
res.encoding = res.apparent_encoding
html_content = res.text.replace('\n', '').replace('\t', '').replace('\r', '')
html_content = re.findall("<script>[\\s.\\S.]*?</script>", html_content)
html_content = html_content[0].replace('<script>', '')
html_content = html_content.replace('</script>', '')
html_content = html_content.replace('icon,', '')
content_list = html_content.split(';')
labels_data = content_list[3].replace(' ', '')
labels_data = labels_data.replace('varLabelsData=', '')
labels_data = labels_data.replace('name', "'name'").replace('position', "'position'").replace('aa', "'aa'").replace('edittime',"'edittime'").replace('href', "'href'").replace(',time', ",'time'")
labels_data = eval(labels_data)

db = pymysql.connect(host='127.0.0.1', user='test',password='123456', database='test', port=3592)
# # 使用cursor()方法获取操作游标
cursor = db.cursor()
print(len(labels_data))
for val in labels_data:
    ref_id = val['href'].replace('/content/?', '').replace('.html', '')
    print(ref_id)
    select_sql = "select id from labels where ref_id = " + ref_id
    # 执行SQL语句
    cursor.execute(select_sql)
    # 获取所有记录列表
    rowcount = cursor.rowcount
    if rowcount > 0:
        continue

    sql = "INSERT INTO labels(`name`, `position`, `status`, `ref_id`)VALUES ('"+str(val['name'])+"', '"+str(val['position'])+"', "+val['aa']+", "+ref_id+")"
    print(sql)
    try:
        # 执行sql语句
        cursor.execute(sql)
        # 提交到数据库执行
        db.commit()
    except:
        # 如果发生错误则回滚
        db.rollback()

# 更新时间
update_time = int(time.time())
update_sql = "update actions set update_time = "+str(update_time)+" where id = 1"
# 执行sql语句
cursor.execute(update_sql)
# 提交到数据库执行
db.commit()

# 关闭数据库连接
db.close()
print("ok")
