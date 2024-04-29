<?php

namespace ppAlgorithm\TwoSum;

class Solution1
{
    /**
     * Найти keys элеиентов из `$nums`, для которых сумма values равна `$target`
     * @param array $nums Массив
     * @param int $target Сумма
     * @return array Массив keys
     */
    public function twoSum(array $nums, int $target): array
    {
        $nums_length = count($nums);
        for ($i = 0; $i < $nums_length; $i++) {
            for ($j = $i + 1; $j < $nums_length; $j++) {
                if ($nums[$i] + $nums[$j] === $target) {
                    return [$i, $j];
                }
            }
        }

        throw new \Exception();
    }
}