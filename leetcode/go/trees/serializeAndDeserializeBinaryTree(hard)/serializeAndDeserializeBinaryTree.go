package main

import (
	"bytes"
	"strconv"
	"strings"
)

/* 297. Serialize and Deserialize Binary Tree
https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/

Сериализация — это процесс преобразования структуры данных или объекта в последовательность битов, чтобы ее можно было
сохранить в файле или буфере памяти или передать по сетевому соединению для последующего восстановления в той же или
другой компьютерной среде.

Разработайте алгоритм сериализации и десериализации двоичного дерева. Нет никаких ограничений на то, как должен работать
ваш алгоритм сериализации/десериализации. Вам просто нужно убедиться, что двоичное дерево может быть сериализовано в
строку, а эта строка может быть десериализована в исходную древовидную структуру.

Пояснение: формат ввода/вывода такой же, как и формат, в котором LeetCode сериализует двоичное дерево. Вам не
обязательно следовать этому формату, поэтому, пожалуйста, проявите креативность и придумайте другие подходы
самостоятельно.

Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
*/

func main() {
	codec := Constructor()

	// Пример дерева: [1,2,3,null,null,4,5]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	// Сериализация
	serialized := codec.serialize(root)
	println("Serialized:", serialized) // "1,2,3,null,null,4,5"

	// Десериализация
	deserialized := codec.deserialize(serialized)

	// Проверка (можно добавить вывод дерева)
	println("Deserialized root:", deserialized.Val) // 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// serialize - Сериализует дерево в одну строку.
func (this *Codec) serialize(root *TreeNode) string {
	var buffer bytes.Buffer

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			buffer.WriteString("N,")
		} else {
			buffer.WriteString(strconv.Itoa(node.Val))
			buffer.WriteString(",")
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)

	return buffer.String()
}

// deserialize - Десериализует закодированные данные в дерево.
func (this *Codec) deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		token := tokens[0]
		tokens = tokens[1:]
		if token == "N" {
			return nil
		}
		val, _ := strconv.Atoi(token)
		return &TreeNode{val, dfs(), dfs()}
	}

	return dfs()
}
