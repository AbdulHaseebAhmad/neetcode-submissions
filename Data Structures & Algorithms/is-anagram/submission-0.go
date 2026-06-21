import "reflect"
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    map1 := make(map[rune]int)
    map2 := make(map[rune]int)

    for index,char := range s {
        value,exists := map1[char]
        
        if !exists {
            map1[char] = 1
        } else {
            map1[char] = value + 1
        }
         r := rune(t[index])

         tvvalue, texists := map2[r]
         if !texists {
            map2[r] = 1
        } else {
            map2[r] = tvvalue + 1
        }
    }
    return reflect.DeepEqual(map1, map2)

}
