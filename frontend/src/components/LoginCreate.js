import { Link } from "react-router-dom";

const LoginCreate = () => {
    return ( 
        <div className="m-10 p-20  " >
            <div className="flex justify-center items-center" >
                <Link to="/login" ><button className=" text-2xl hover:text-white m-2 p-2 bg-yellow-500 rounded-xl shadow-xl " >Login</button></Link>
                <Link to="/signup" ><button className=" text-2xl hover:text-white m-2 p-2 bg-yellow-500 rounded-xl shadow-xl " >Sign Up</button></Link>
            </div>
        </div>
     );
}
 
export default LoginCreate;