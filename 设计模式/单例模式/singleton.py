#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   singleton.py
@Time    :   2022/09/12 21:13:18
@Author  :   scliang 
@Desc    :   单例模式
'''

'''
保证一个类只有一个实例，并提供一个放问它的全局访问点.
角色：
    单例
优点： 
    1. 对唯一实例的受控访问.
    2. 单例相当于全局变量，但防止了命名空间被污染.
'''


# here put the import lib

class SingletonPattern():
    def __new__(cls: type[Self]) -> Self:
        if not hasattr(cls,"_instance"):
            cls._instance = super(SingletonPattern,cls).__new__(cls)
        return cls._instance

class MyClass(SingletonPattern):
    def __init__(self,number) -> None:
        self.number = number


# client
n1 = MyClass(100)
n2 = MyClass(200)

# n1,n2 是返回的同一个实例，但是n2后修改，所以修改了全局实例对象的值
print(n1) # 200
print(n2) # 200

# 查看是否是用一个对象
print(id(n1) == id(n2))  # True