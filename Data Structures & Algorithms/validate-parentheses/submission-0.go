func isValid(s string) bool {

    if len(s)%2 != 0 {
        return false
    }

    pa := map[rune]rune{
        '}': '{',
        ')': '(',
        ']': '[',
    }

    var check []rune  // ✅ slice, actúa como pila

    for _, v := range s {
        if v == '{' || v == '(' || v == '[' {
            check = append(check, v)  // push
            continue
        }
        if v == '}' || v == ')' || v == ']' {
            if len(check) == 0 || check[len(check)-1] != pa[v] {
                return false
            }
            check = check[:len(check)-1]  // pop
        }
    }

    return len(check) == 0
}