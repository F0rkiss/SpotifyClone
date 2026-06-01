import {  EyeAlt, EyeClosed} from "@boxicons/react"
import { useState } from "react"

type FormInputProps ={
    id : string
    label : string
    type? : string
    placeholder? : string
}

const FormInput = ({id, label, type="text", placeholder}: FormInputProps) => {
    const [showPassword, setShowPassword] = useState(false)
    
    const isPassword = type ==="password"
    const inputType = isPassword && showPassword ? "text" : type

    

  return (
    <div className="flex flex-col gap-2 pb-3 w-80">
        <label htmlFor={id}>{label}</label>
        <div className="flex rounded border 
            border-neutral-700
            hover:border-neutral-500
            focus-within:border-neutral-200 
            bg-black px-3 py-2 text-white ">
            <input 
            id={id}
            type={inputType} 
            placeholder={placeholder}
            className="w-full"
            >
            </input>
            {isPassword && (
            <div className="text-right">
                <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="flex-col flex items-center pl-2">
                    {showPassword ? <EyeClosed /> : <EyeAlt />}
                </button>
            </div>)}
        </div>
    </div>
  )
}

export default FormInput