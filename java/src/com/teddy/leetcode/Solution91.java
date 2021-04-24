package com.teddy.leetcode;

import java.util.Arrays;

/**
 * @link https://leetcode-cn.com/problems/decode-ways/
 */
class Solution91 {

    public static void main(String[] args) {
        String s = "2611055971756562";
        System.out.println(new Solution91().numDecodings(s));
    }

    public int numDecodings(String s) {
        int[] memo = new int[s.length()];
        Arrays.fill(memo, -1);
        return countNum(s, 0, memo);
    }

    /**
     * 递归
     *
     * @param s     输入字符串
     * @param start 下标
     * @return 字符串从下标开始的解码数量
     */
    private int countNum(String s, int start, int[] memo) {
        if (start >= s.length()) {
            return 1;
        }

        if (memo[start] > -1) {
            return memo[start];
        }

        int doubleCount = 0;
        int singleCount = 0;

        // double chars
        if (start + 1 < s.length()) {
            if (s.charAt(start) >= '1' && s.charAt(start) <= '2' &&
                    // could not gt 26
                    (s.charAt(start + 1) <= '6') || s.charAt(start) == '1') {
                doubleCount = countNum(s, start + 2, memo);
            }
        }

        // single char
        if (s.charAt(start) > '0') {
            singleCount = countNum(s, start + 1, memo);
        }

        int result = doubleCount + singleCount;
        memo[start] = result;
        return result;
    }
}