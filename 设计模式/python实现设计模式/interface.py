#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   interface.py
@Time    :   2022/09/12 21:21:05
@Author  :   scliang 
@Desc    :   None
'''

# here put the import lib
# 引入抽象类
from abc import ABCMeta, abstractmethod
import re

class Payment(metaclass=ABCMeta):
    # abstract class
    @abstractmethod
    def pay(self,money):
        pass


# 实现抽象类
class AliPay(Payment):
    def pay(self,money):
       print("支付宝支付%d" %money)

class WechatPay(Payment):
    def pay(self, money):
        print("微信支付%d元" % money)



# 直接实例化一个抽象类会直接报错
# p = Payment()
# print("abstract method:",p)
'''
Traceback (most recent call last):
  File "F:\GoProjects\src\MyPractiseNotes\设计模式\interface.py", line 33, in <module>
    p = Payment()
TypeError: Can't instantiate abstract class Payment with abstract method pay
'''

# 可以实例化抽象类的实现类
p = AliPay()
p.pay(100)

"""
PS F:\GoProjects\src\MyPractiseNotes\设计模式> python .\interface.py
支付宝支付100
"""