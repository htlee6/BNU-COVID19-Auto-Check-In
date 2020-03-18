# -*- coding: utf-8 -*-
# /usr/bin/python
import requests, json, re
import time, datetime, sys


login_url = "https://app.bupt.edu.cn/uc/wap/login?redirect=https%3A%2F%2Fapp.bupt.edu.cn%2Fncov%2Fwap%2Fdefault%2Findex"
base_url = "https://app.bupt.edu.cn/ncov/wap/default/index"
save_url = "https://app.bupt.edu.cn/ncov/wap/default/save"
login_check_url = "https://app.bupt.edu.cn/uc/wap/login/check"

class DaKa(object):
    def __init__(self, username, password):
        self.username = username
        self.password = password
        self.name = ""
        self.number = ""
        self.sess = requests.Session()

    def login(self):
        """Login to BUPT platform"""
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
        old_info = json.loads(re.findall(r'oldInfo: (.*)', html)[0][:-1])
        name = re.findall(r'realname: "([^\"]+)",', html)[0]
        number = re.findall(r"number: '([^\']+)',", html)[0]

        new_info = old_info.copy()
        new_info['name'] = name
        new_info['number'] = number
        new_info["date"] = self.get_date()
        new_info["created"] = round(time.time())

        self.info = new_info
        return new_info

    def _rsa_encrypt(self, password_str, e_str, M_str):
        password_bytes = bytes(password_str, 'ascii') 
        password_int = int.from_bytes(password_bytes, 'big')
        e_int = int(e_str, 16) 
        M_int = int(M_str, 16) 
        result_int = pow(password_int, e_int, M_int) 
        return hex(result_int)[2:].rjust(128, '0')

    def daka(self):
        res = self.login()
        print("--------login res-------\n", res)
        res = self.get_info()
        print("--------get_info res-------\n", res)
        res = self.post()
        print("--------post res-------\n", res)
        return res



def main(username, password):

    dk = DaKa(username, password)
    try:
        res = dk.daka()

        if res['e'] == 0:
            print("打卡成功")
        else:
            print("打卡失败，原因：{}".format(res["m"]))
    except Exception as e:
        print(e)
    finally:
        return 0



if __name__=="__main__":
    #username = sys.argv[1]
    #password = sys.argv[2]
    main("2019110680","2019110680tch...")
