#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   chain_responsibility.py
@Time    :   2022/09/13 14:36:00
@Author  :   scliang 
@Desc    :   责任链模式
'''

'''
使多个对象都有机会处理请求，从而避免请求的发送者和接受者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理该请求为止.
角色：
    抽象处理者
    具体处理者
    客户端
优点：
    降低耦合，一个对象不需要知道是其他那个兑现处理请求.
'''

# here put the import lib
from abc import ABCMeta, abstractmethod

# 抽象的处理者
class Handler(metaclass=ABCMeta):
    @abstractmethod
    def handle_leave(self, day):
        pass

# 具体的处理者
class GeneralManager(Handler):
    def handle_leave(self, day):
        if day <= 10:
            print('总经理准假%d天' % day)
        else:
            print('可以辞职了！')

# 具体的处理者
class DepartmentManager(Handler):
    def __init__(self):
        self.next = GeneralManager()

    def handle_leave(self, day):
        if day <= 7:
            print('项目主管准假%d天' % day)
        else:
            print('部门经理职权不足')
            self.next.handle_leave(day)

# 具体的处理者
class ProjectDirector(Handler):
    def __init__(self):
        self.next = DepartmentManager()

    def handle_leave(self, day):
        if day <= 3:
            print('项目主管准假%d天' % day)
        else:
            print('项目主管职权不足')
            self.next.handle_leave(day)

day = 10
p = ProjectDirector()
p.handle_leave(day)
"""
out:
    项目主管职权不足
    部门经理职权不足
    总经理准假10
"""