# 347. Top K Frequent Elements

## Situation

Given an integer array `nums` and an integer `k`, return *the* `k` *most frequent elements*. You may return the answer in **any order**.

**Example 1:**

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is **guaranteed** that the answer is **unique**.

 

**Follow up:** Your algorithm's time complexity must be better than `O(n log n)`, where n is the array's size.

## Условие задачи

Дан *integer array* `nums` и *integer* `k`. Нужно написать функцию, которая возвращает `k` наиболее часто встречающихся *element*'ов в *array*. Вы можете вернуть ответ в **любом порядке**.

**Пример 1:**

```
Input: nums = [1, 1, 1, 2, 2, 3], k = 2
Output: [1, 2]
```

**Пример 2**:

```
Input: nums = [1], k = 1
Output: [1]
```

**Ограничения:**

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `k` находится в диапазоне `[1, число уникальных element'ов в array]`.
- Гарантируется **,** что ответ **единственный** (**unique**) .

 

*Time complexity* вашего алгоритма должна быть лучше, чем `O(n log n)`, где `n` — *array size*.



## Решение 1.  Hash table + сортировка

1. Считаем, сколько раз встречается каждый из *element*'ов в *array* (`count`). Для этого используем *hash table* (храним `element: count`)
2. Создаем и заполняем *array* из уникальных исходных *element*'ов (или из пар `{element, count}`, но так требуется больше памяти).
3. Сортируем получившийся *array* по убыванию по значению `count`
4. Выбираем из отсортированного *array* `k` первых *element*'ов.

[Код](solution1.go)

```go
func top1(nums []int, k int) []int {
	counters := make(map[int]int)

	for _, num := range nums {
		counters[num]++
	}

	uniqueNums := make([]int, 0, len(counters))
	for num := range counters {
		uniqueNums = append(uniqueNums, num)
	}

	sort.Slice(uniqueNums, func(i, j int) bool {
		return counters[uniqueNums[i]] > counters[uniqueNums[j]]
	})

	if k > len(uniqueNums) {
		return uniqueNums
	}

	return uniqueNums[:k]
}
```

[Тесты](solution_test.go)

**Time complexity**. `O(N⋅logN))` – где `N` – количество уникальных *element*'ов в *array* (в худшем случае равно общему количеству *element*'ов).

## Решение 2. Heap







