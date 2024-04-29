<?php

namespace ppAlgorithm\TwoSum;

use Exception;

class Solution3
{
    /**
     * Найти keys элеиентов из `$nums`, для которых сумма values равна `$target`
     * @param array $nums Массив
     * @param int $target Сумма
     * @return array Массив keys
     * @throws \Exception
     */
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

        throw new Exception();
    }
}