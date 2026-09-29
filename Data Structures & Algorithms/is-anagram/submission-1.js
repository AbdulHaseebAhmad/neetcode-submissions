class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
        const keycount = new Map();
        if (s.length !== t.length) {
            return false;
        }
        s = s.split("").sort().join("");
        t = t.split("").sort().join("");

        for (let i = 0; i < s.length; i++) {
            keycount.set(s[i], keycount.has(s[i]) ? keycount.get(s[i]) + 1 : 1);

            if (!keycount.has(t[i])) {
                return false;
            }
            keycount.set(t[i], keycount.get(t[i]) - 1);

        }

         for (const count of keycount.values()) {
        if (count !== 0) {
            return false;
        }
    }

    return true;
    }
}
