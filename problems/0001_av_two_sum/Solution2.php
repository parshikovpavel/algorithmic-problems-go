<?php

namespace ppAlgorithm\TwoSum;

use Exception;

class Solution2
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
            $hashtable[$value] = $key;
        }

        foreach ($nums as $key => $value) {
            $complement = $target - $value;
            if (isset($hashtable[$complement]) && $hashtable[$complement] !== $key) {
                return [$key, $hashtable[$complement]];
            }
        }

        throw new Exception();
    }
}