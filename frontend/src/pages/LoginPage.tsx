import GreenButton from "../components/ui/GreenButton"
import WhiteLogo from "../assets/spotify-logo/white_logo.png"
import FormInput from "../components/ui/FormInput"

const LoginPage = () => {
  return (
    <div className='flex flex-col min-h-screen justify-center items-center'>
        <div className="flex">
            <img className="w-12 " src={WhiteLogo} alt="WhiteLogo" />
        </div>
        <div className="flex bold pt-4 pb-8 text-5xl font-bold">
            Welcome back
        </div>
        <FormInput id="identifier" label="Username or Email"/>
        <FormInput id="password" label="password" type="password"/>
        <GreenButton label="belom" />
    </div>
  )
}

export default LoginPage