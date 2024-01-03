---
created: 2024-01-03T17:44:56 (UTC +08:00)
tags: []
source: https://leetcode.cn/problems/remove-nodes-from-linked-list/description/
author: 
---

# 2487. 从链表中移除节点 - 力扣（LeetCode）

给你一个链表的头节点 head 。

移除每个右侧有一个更大数值的节点。

返回修改后链表的头节点 head 。


示例 1：

[https://assets.leetcode.com/uploads/2022/10/02/drawio.png]


输入：head = [5,2,13,3,8]
输出：[13,8]
解释：需要移除的节点是 5 ，2 和 3 。
- 节点 13 在节点 5 右侧。
- 节点 13 在节点 2 右侧。
- 节点 8 在节点 3 右侧。


示例 2：


输入：head = [1,1,1,1]
输出：[1,1,1,1]
解释：每个节点的值都是 1 ，所以没有需要移除的节点。


 

提示：

* 给定列表中的节点数目在范围 [1, 105] 内
* 1 <= Node.val <= 105

---
给你一个链表的头节点 `head` 。

移除每个右侧有一个更大数值的节点。

返回修改后链表的头节点 `head` 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2022/10/02/drawio.png)

```
<strong>输入：</strong>head = [5,2,13,3,8]
<strong>输出：</strong>[13,8]
<strong>解释：</strong>需要移除的节点是 5 ，2 和 3 。
- 节点 13 在节点 5 右侧。
- 节点 13 在节点 2 右侧。
- 节点 8 在节点 3 右侧。
```

**示例 2：**

```
<strong>输入：</strong>head = [1,1,1,1]
<strong>输出：</strong>[1,1,1,1]
<strong>解释：</strong>每个节点的值都是 1 ，所以没有需要移除的节点。
```

**提示：**

-   给定列表中的节点数目在范围 `[1, 10<sup>5</sup>]` 内
-   `1 <= Node.val <= 10<sup>5</sup>`

