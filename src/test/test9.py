from pymysqlreplication import BinLogStreamReader

mysql_settings = {'host': '116.62.213.56', 'port': 51603, 'user': 'root', 'passwd': 'huansi@2017'}
#mysql -h 116.62.213.56 -P 51603 -u root -phuansi@2017 

stream = BinLogStreamReader(connection_settings = mysql_settings, server_id=100)

for binlogevent in stream:
    binlogevent.dump()

stream.close()
