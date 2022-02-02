# [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)

## Situation

You are given two integer arrays `nums1` and `nums2`, sorted in **non-decreasing order**, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively.

**Merge** `nums1` and `nums2` into a single array sorted in **non-decreasing order**.

The final sorted array should not be returned by the function, but instead be *stored inside the array* `nums1`. To accommodate this, `nums1` has a length of `m + n`, where the first `m` elements denote the elements that should be merged, and the last `n` elements are set to `0` and should be ignored. `nums2` has a length of `n`.

 

**Example 1:**

```
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
```

**Example 2:**

```
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
```

**Example 3:**

```
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
```

 

**Constraints:**

- `nums1.length == m + n`
- `nums2.length == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-109 <= nums1[i], nums2[j] <= 109`

 

**Follow up:** Can you come up with an algorithm that runs in `O(m + n)` time?



## Условие задачи

Вам даны два *array*'s из *integer*'s `nums1` и `nums2`, отсортированные в **неубывающем порядке** (возрастающем), и два *integer*'s `m` и `n`, представляющие количество элементов в `nums1`и `nums2` соответственно.

**Смержить** `nums1` и `nums2` в один *array*, отсортированный в **неубывающем порядке** (возрастающем).

Окончательный отсортированный *array* не должен возвращаться функцией, а должен *храниться внутри array* `nums1` . Чтобы приспособиться к этому, `nums1` имеет *length* `m + n`, где первые `m` элементов обозначают элементы, которые должны быть смержены, а последние `n` элементов установлены в `0` и должны игнорироваться. `nums2` имеет *length* `n`.

 **Пример 1:**

```
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Объяснение: Были смержены array's [1,2,3] и [2,5,6].
Результат мержа [1,2,2,3,5,6] с underlined array `nums1`.
```

**Пример 2:**

```
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Объяснение: Были смержены array's [1] и [].
Результат мержа – [1].
```

**Пример 3:**

```
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Объяснение: Были смержены array's – [] и [1].
Результат мержа – [1].
Обратите внимание, что поскольку m = 0, в nums1 нет элементов. 0 нужен только для того, чтобы результат мержа поместился в nums1.
```

 

**Ограничения:**

- `nums1.length == m + n`
- `nums2.length == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-10^9 <= nums1[i], nums2[j] <= 10^9`

 

**Follow up:** можете ли вы придумать алгоритм, работающий за `O(m + n)` *time*?



## Решение

- Т.к. в  конце массива `a` есть пустое место для хранения элементов, чтобы смержить массивы *in place* – начинаем сравнивать элементы массивов с конца - в начало (от больших к меньшим)
- заводим три указателя – `aI`, `bI`, `resI` для прохода по части массива `a` (где есть числа, размером `m`), массиву `b` и всему массиву `a` с результатом.
- На каждом шаге выбираем наибольший элемент из конца массивов `a` (той части где есть числа, размером `m`) и `b` и пишем эти элементы в конец массива `a` (полного массива, где будет весь результат)
- Если массив `b` закончится первым, то оставшиеся элементы массива `a` уже будут стоят на своих местах
- Если закончится первым массив `a`, то надо скопировать оставшиеся элементы из массива `b` в начало массива `a`.

[Код](solution.go)

```go
func merge(a []int, m int, b []int, n int) {
	aI := m - 1
	bI := n - 1
	resI := m + n - 1

	for aI >= 0 && bI >= 0 {
		if a[aI] > b[bI] {
			a[resI] = a[aI]
			aI--
		} else {
			a[resI] = b[bI]
			bI--
		}
		resI--
	}

	copy(a[:resI+1], b[:bI+1])
}
```

[Тесты](solution_test.go)

