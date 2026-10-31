package validator

import (
	"fmt"
	"reflect"
	"strings"
)

func ValidateNotEmpty(values ...interface{}) error {
    for i, val := range values {
        v := reflect.ValueOf(val)
        isEmpty := false

        switch v.Kind() {
        case reflect.Invalid: 
            isEmpty = true
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            isEmpty = v.Int() == 0
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            isEmpty = v.Uint() == 0
        case reflect.Float32, reflect.Float64:
            isEmpty = v.Float() == 0
        case reflect.Complex64, reflect.Complex128:
            isEmpty = v.Complex() == 0
        case reflect.String:
            isEmpty = v.Len() == 0
        case reflect.Slice, reflect.Array, reflect.Map:
            isEmpty = v.Len() == 0
        case reflect.Ptr, reflect.Interface:
            isEmpty = v.IsNil()
        default:            
            continue
        }

        if isEmpty {
            return fmt.Errorf("value at index %d (type %s) must not be empty", i, v.Type())
        }
    }
    return nil
}

func ValidateNonNegative(values ...interface{}) error {
    for i, val := range values {
        v := reflect.ValueOf(val)
        var isNegative bool

        switch v.Kind() {
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            isNegative = v.Int() < 0
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            continue
        case reflect.Float32, reflect.Float64:
            isNegative = v.Float() < 0
        default:
            return fmt.Errorf("value at index %d (type %s) is not a numeric type", i, v.Type())
        }

        if isNegative {
            return fmt.Errorf("value at index %d (type %s) must be non-negative (got %v)", i, v.Type(), val)
        }
    }
    return nil
}

func ValidatePositive(values ...interface{}) error {
    for i, val := range values {
        v := reflect.ValueOf(val)
        var isNotPositive bool

        switch v.Kind() {
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            isNotPositive = v.Int() <= 0
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            isNotPositive = v.Uint() == 0
        case reflect.Float32, reflect.Float64:
            isNotPositive = v.Float() <= 0
        default:
            return fmt.Errorf("value at index %d (type %s) is not a numeric type", i, v.Type())
        }

        if isNotPositive {
            return fmt.Errorf("value at index %d (type %s) must be positive (got %v)", i, v.Type(), val)
        }
    }
    return nil
}

func ValidateMin(min float64, values ...interface{}) error {
    for i, val := range values {
        v := reflect.ValueOf(val)
        var actual float64
        switch v.Kind() {
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            actual = float64(v.Int())
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            actual = float64(v.Uint())
        case reflect.Float32, reflect.Float64:
            actual = v.Float()
        default:
            return fmt.Errorf("value at index %d (type %s) is not numeric", i, v.Type())
        }
        if actual < min {
            return fmt.Errorf("value at index %d (type %s) must be >= %v (got %v)", i, v.Type(), min, val)
        }
    }
    return nil
}

func ValidateMax(max float64, values ...interface{}) error {
    for i, val := range values {
        v := reflect.ValueOf(val)
        var actual float64
        switch v.Kind() {
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            actual = float64(v.Int())
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            actual = float64(v.Uint())
        case reflect.Float32, reflect.Float64:
            actual = v.Float()
        default:
            return fmt.Errorf("value at index %d (type %s) is not numeric", i, v.Type())
        }
        if actual > max {
            return fmt.Errorf("value at index %d (type %s) must be <= %v (got %v)", i, v.Type(), max, val)
        }
    }
    return nil
}

func ValidateRange(min, max float64, values ...interface{}) error {
    if err := ValidateMin(min, values...); err != nil {
        return err
    }
    return ValidateMax(max, values...)
}

func ValidateLength(length int, values ...interface{}) error {
    for i, val := range values {
        v := reflect.ValueOf(val)
        kind := v.Kind()
        if kind != reflect.String && kind != reflect.Slice && kind != reflect.Array && kind != reflect.Map {
            return fmt.Errorf("value at index %d (type %s) does not have length", i, v.Type())
        }
        if v.Len() != length {
            return fmt.Errorf("value at index %d (type %s) must have length %d (got %d)", i, v.Type(), length, v.Len())
        }
    }
    return nil
}

func ValidateEmail(values ...string) error {
    for i, email := range values {
        if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
            return fmt.Errorf("value at index %d is not a valid email: %s", i, email)
        }
    }
    return nil
}