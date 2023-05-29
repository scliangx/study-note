#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   proxy.py
@Time    :   2022/09/13 14:03:54
@Author  :   scliang 
@Desc    :   代理模式
'''

'''
为其它对象提供一种代理以控制对这个对象的访问.
角色：
    抽象实体
    实体
    代理
优点：
    1. 远程代理：可以隐藏对象位于远程地址空间的事实
    2. 虚代理：可以进行优化，例如根据要求创建对象
    3. 保护代理：允许访问一个对象时附加一些内务处理
'''

# here put the import lib
from abc import ABCMeta, abstractmethod

class Subject(metaclass=ABCMeta):
    @abstractmethod
    def get_content(self):
        pass

    @abstractmethod
    def set_content(self, content):
        pass

# 远程代理
class RealSubject(Subject):
    def __init__(self, filename):
        self.filename = filename
        print('读取文件内容！')
        with open(self.filename, 'r', encoding='utf-8') as f:
            self.content = f.read()

    def get_content(self):
        return self.content

    def set_content(self, content):
        with open(self.filename, 'w', encoding='utf-8') as f:
            f.write(content)


subj = RealSubject('test.txt')
'''
out: 读取文件内容！
'''



# 虚代理
# 只有使用的时候才会实例化RealSubject(self.filename) 对象
class VirtualProxy(Subject):
    def __init__(self, filename):
        self.filename = filename
        self.subj = None

    def get_content(self):
        if not self.subj:
            self.subj = RealSubject(self.filename)
        return self.subj.get_content()

    def set_content(self, content):
        if not self.subj:
            self.subj = RealSubject(self.filename)

        return self.subj.set_content(content)

subj = VirtualProxy('test.txt')
print(subj.get_content())
'''
out:
    读取文件内容！
    testting....
'''


# 保护代理
class ProtectedSubject(Subject):
    def __init__(self, filename):
        self.subj = RealSubject(filename)

    def get_content(self):
        return self.subj.get_content()

    def set_content(self, content):
        raise PermissionError('无写入权限！')

subj = ProtectedSubject('test.txt')
print(subj.get_content())
subj.set_content('abc')

'''
out:
    取文件内容！
    testting....
    Traceback (most recent call last):
    File "F:\GoProjects\src\MyPractiseNotes\设计模式\代理模式\proxy.py", line 89, in <module>
        subj.set_content('abc')
    File "F:\GoProjects\src\MyPractiseNotes\设计模式\代理模式\proxy.py", line 85, in set_content
        raise PermissionError('无写入权限！')
    PermissionError: 无写入权限！
'''