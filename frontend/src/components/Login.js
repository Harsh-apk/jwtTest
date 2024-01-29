import { useState } from "react";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";


const Login = () => {
    const historyStack = useHistory();
    const [email,setEmail] = useState("");
    const [password,setPassword] = useState("");
    const [error,setError] = useState(null);
    const handleSubmit = async(e)=>{
        setError(null);
        e.preventDefault();
        const body = {email,password};
        try{
            const res  = await fetch ("http://127.0.0.1:5000/login",{
            method:'POST',
            mode:'cors',
            headers:{
                "Content-Type":"application/json"
            },
            body:JSON.stringify(body),
            redirect:"follow"
        });
        if(!res.ok){
            throw new Error("Something went wrong 😢")
        }
        const response = await res.json();
        if(response.error){
            throw new Error(response.error)
        }
        //setUser(response);
        historyStack.push("/");
        return;

        }catch (error){
             setError(error.message);
             return;
        }

    }
    return ( 
        <div className="m-10" >
            {error && <div className=" text-center text-red-600 " >{error}</div>}
            <form onSubmit={handleSubmit}>
                <div className="m-2 p-2 font-bold text-xl text-center" ><label>Email</label></div>
                <div className="flex justify-center" ><input className="m-2 p-2  border border-yellow-500 rounded-xl text-center " type="email" required value={email} onChange={(e)=>{setEmail(e.target.value)}} /></div>
                <div className="m-2 p-2 font-bold text-xl text-center"><label>Password</label></div>
                <div className=" flex justify-center " ><input className="m-2 p-2 border border-yellow-500 rounded-xl text-center " type="password" required value={password} onChange={(e)=>{setPassword(e.target.value)}} /></div>
                <div className="flex justify-center flex-wrap m-2 p-2 " ><button className="text-xl font-bold shadow-xl  hover:text-white bg-yellow-500 p-2 rounded-xl " > Login </button></div>
            </form>
        </div>
     );
}
 
export default Login;