# Two sum (AV)

## Условие

Дан список целых неповторяющихся чисел и целое число. Необходимо найти все пары в списке, сумма которых равна заданному числу. Перестановки `(i, j)` и `(j, i)` считаются одной и той же парой, нужно вывести только одну из них (любую).

**Input**

- массив целых чисел (`[]int`) размера `N`. 

- целое число.

**Output**

- список пар.

**Example 1:**

```csharp
Input: [2, 4, 5, 3] 7
Output: [ [2, 5], [4, 3] ]
```

**Example 2:**

```
Input: [-5, 2, 12, 5] 7
Output: [ [-5, 12], [2, 5] ]
```



| #    | Time               | Space                    |
| ---- | ------------------ | ------------------------ |
| 1    | *O(N<sup>2</sup>)* | *O(1)*                   |
| 2    | *O(N)*             | *O(N)* (для *Hashtable*) |
| 3    | *O(N)*             | *O(N)* (для *Hashtable*) |

## Brute-force search

Самое простое решение в лоб – *Brute-force search* ([1](https://github.com/parshikovpavel/cheat-sheets/blob/master/Algorithm.md#brute-force-search))

Проходимся по всем элементам `$nums[$i]` массива и пытаемся найти такой `$nums[$j]`, для которого `$nums[$i] + $nums[$j] === $target`. Достаточно перебирать только `$j > $i`. В этом случае перестановки проверяем единожды - вместо `(i, j)` и `(j, i)` - проверяем только `(i, j)`.

[Реализация](Solution1.php):

```go
func twoSum(nums []int, target int) (result [][2]int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); i++ {
			if nums[i]+nums[j] == target {
				result = append(result, [2]int{nums[i], nums[j]})
			}
		}
	}

	return
}

```

[Тесты](./../../tests/TwoSum/Solution1Test.php)

TODO!!!



## Два прохода с Hashtable

Алгоритм:

- Построим *Hashtable* по *value*'s из `$nums`.  Это позволит сократить время поиска с *O(n)* до *O(1)*, при этом необходимо затратить *O(n)* памяти под *Hashtable* (*space-time trade-off*). 
- Проходим по `$nums` и пытаемся найти в *Hashtable* такое *value*, которое дополняет `$nums[$i]` до `$target`, т.е. равное `$target - $nums[$i]`. Необходимо добавить проверку, что `$i !== $j`.

[Реализация](Solution2.php):

```php
public function twoSum(array $nums, int $target): array
{
    $hashtable = [];

    foreach ($nums as $key => $value) {
        $hashtable[$value] = $key;
    }

    foreach ($nums as $key => $value) {
        $complement = $target - $value;
        if (isset($hashtable[$complement]) && $hashtable[$complement] !== $key) {
            return [$key, $hashtable[$complement]];
        }
    }
}
```

[Тесты](./../../tests/TwoSum/Solution2Test.php)

## Один проход с Hashtable

Алгоритм:

- Так же как и в предыдущем случае, проходим по `$nums` и строим *Hashtable* по *value*'s. 
- На каждой итерации, не дожидаясь окончания цикла, сразу же пытаемся найти в *Hashtable* такое *value*, которое дополняет `$nums[$i]` до `$target`, т.е. равное `$target - $nums[$i]`. Т.к. этот поиск выполняется до вставки текущего *value* в *Hashtable*, проверка что `$i !== $j` не требуется.

[Реализация](Solution3.php):

```php
public function twoSum(array $nums, int $target): array
{
    $hashtable = [];

    foreach ($nums as $key => $value) {
        $complement = $target - $value;
        if (isset($hashtable[$complement])) {
            return [$hashtable[$complement], $key];
        }
        $hashtable[$value] = $key;
    }

    throw new \Exception();
}
```

[Тесты](./../../tests/TwoSum/Solution3Test.php)

