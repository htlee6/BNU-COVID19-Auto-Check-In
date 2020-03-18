# -*- coding: utf-8 -*-
# /usr/bin/python
import requests, json, re
import time, datetime, sys


login_url = "https://app.buaa.edu.cn/uc/wap/login?redirect=https%3A%2F%2Fapp.buaa.edu.cn%2Fncov%2Fwap%2Fdefault%2Findex%3Ffrom%3Dhistory"
base_url = "https://app.buaa.edu.cn/ncov/wap/default/index"
save_url = "https://app.buaa.edu.cn/ncov/wap/default/save"
login_check_url = "https://app.buaa.edu.cn/uc/wap/login/check"

class DaKa(object):
    def __init__(self, username, password):
        self.username = username
        self.password = password
        self.name = ""
        self.number = ""
        self.sess = requests.Session()

    def login(self):
        """Login to BUAA platform"""
        res = self.sess.get(login_url)
        
        if res.status_code != 200:
            raise Exception("{} 登陆平台失败，失败代码{}".format(self.username,res.status_code))
        #n, e = res['modulus'], res['exponent']
        #encrypt_password = self._rsa_encrypt(self.password, e, n)

        data = {
            'username': self.username,
            'password': self.password,
        }
        res = self.sess.post(url=login_check_url, data=data)
        ret = json.loads(res.content.decode())

        if ret['e'] != 0:
            raise Exception("{} 登陆失败，原因{}".format(self.username,ret['m']))
        return ret
    
    def post(self):
        """Post the hitcard info"""
        res = self.sess.post(save_url, data=self.info)
        if res.status_code != 200:
            raise Exception("{} post info faild, statu code = {}".format(self.username,res.status_code))
        return json.loads(res.text)
    
    def get_date(self):
        today = datetime.date.today()
        return "%4d%02d%02d" %(today.year, today.month, today.day)
        
    def get_info(self, html=None):
        """Get hitcard info, which is the old info with updated new time."""
        if not html:
            res = self.sess.get(base_url)
            if res.status_code != 200:
                raise Exception("{} get info faild, statu code = {}".format(self.username,res.status_code))
            html = res.content.decode()
        
        old_info_raw = re.findall(r'oldInfo: (.*)', html)[0][:-1]

        # 解决 geo_api_info 字段奇怪的转义问题
        geo_info_raw = re.findall(r'"geo_api_info":"(.*)","area"', old_info_raw)[0]
        geo_info = json.loads(geo_info_raw.replace('\\"','"'))
        geo_info_match = re.search(r'("geo_api_info":.*",)"area"', old_info_raw)
        geo_regs = geo_info_match.regs[1]

        old_info_raw = old_info_raw[0:geo_regs[0]]  + old_info_raw[geo_regs[1]:]
        old_info = json.loads(old_info_raw)

        name = re.findall(r'realname: "([^\"]+)",', html)[0]
        number = re.findall(r"number: '([^\']+)',", html)[0]

        self.name = name
        self.number = number

        new_info = old_info.copy()
        new_info.update({"geo_api_info":geo_info})
        new_info["date"] = self.get_date()
        new_info["created"] = round(time.time())
        
        # 未知的奇怪参数
        new_info.update({"gwszdd":""})
        new_info.update({"sfyqjzgc":""})
        new_info.update({"jrsfqzys":""})
        new_info.update({"jrsfqzfy":""})
        self.info = new_info

        self.realname = name
        self.school_id = number
        self.position = geo_info['formattedAddress']
        return new_info

    def _rsa_encrypt(self, password_str, e_str, M_str):
        password_bytes = bytes(password_str, 'ascii') 
        password_int = int.from_bytes(password_bytes, 'big')
        e_int = int(e_str, 16) 
        M_int = int(M_str, 16) 
        result_int = pow(password_int, e_int, M_int) 
        return hex(result_int)[2:].rjust(128, '0')

    def daka(self):
        self.login()
        self.get_info()
        ret =  self.post()
        return ret



def main(username, password):

    dk = DaKa(username, password)
    try:
        res = dk.daka()
        
        if ret['e'] == 0:
            print("打卡成功")
        else:
            print("打卡失败，原因：{}".format(ret["m"]))
    except Exception as e:
        print(e)
    finally:
        return 0



if __name__=="__main__":
    #username = sys.argv[1]
    #password = sys.argv[2]
    main("fjlsso","Fzw20859")
