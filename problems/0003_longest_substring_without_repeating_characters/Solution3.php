<?php

namespace ppAlgorithm\LongestSubstringWithoutRepeatingCharacters;

class Solution3
{
    /**
     * Найти максимальную длину подстроки без повторения символов
     * @param string $s
     * @return int
     */
    public function lengthOfLongestSubstring(string $s): int
    {
        $i = 0;
        $hashTable = [];
        $lengthS = strlen($s);
        $maxLength = 0;

        for ($j = 0; $j < $lengthS; $j++) {
            if (isset($hashTable[$s[$j]]) && $hashTable[$s[$j]] >= $i) {
                $i = $hashTable[$s[$j]] + 1;
            }
            else {
                $maxLength = max($j - $i + 1, $maxLength);
            }
            $hashTable[$s[$j]] = $j;
        }

        return $maxLength;
    }
}