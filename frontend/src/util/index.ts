export const Sleep = (ms: number) => {
    return new Promise(resolve => setTimeout(resolve, ms))
}

export const SUCESS_CODE = "100"
export const FAIL_CODE = "500"