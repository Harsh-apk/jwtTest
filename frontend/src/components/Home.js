import { useEffect, useState } from "react";
import { useHistory } from "react-router-dom";

const Home = ({user, setUser}) => {
  const historyStack = useHistory();
  const [pending, setPending] = useState(false);
  const [error, setError] = useState(null);
    useEffect(()=>{
      const getData = async() => {
        setPending(true);
        setError(null);
        try {
          const res = await fetch("http://127.0.0.1:5000/api/v1/user", {
            method: "GET",
            headers: {
              Accept: "application/json",
              "Content-Type": "application/json",
            },
            credentials: "include",
            
          });
          if (!res.ok) {
            throw new Error("Something went wrong 🥲");
          }
          const response = await res.json();
          if (response.error) {
            throw new Error("Sommething went wrong 🥲");
          }
          setUser(response);
        } catch (err) {
          setPending(false);
          setError(err.message);
        }finally{
          setPending(false)
        }
      }; 
      if(user===null){
        getData();
      }
      if (error) {
        historyStack.push("/loginCreate");
      }
    })
    
  return (
    <div className="m-10">
      {error && (
        <div className=" text-red-500 text-center m-2 p-2 ">{error}</div>
      )}
      {pending && (
        <div className=" text-red-500 text-center m-2 p-2 ">Pending...</div>
      )}
      {user && (
        <div className="m-5 p-5 rounded-3xl shadow-2xl ">
          <div className="p-5 m-2" >Id : {user.id}</div>
          <div className="p-5 m-2"  >Email : {user.email}</div>
          <div className="p-5 m-2"  >User Name : {user.userName}</div>
        </div>
      )}
    </div>
  );
};

export default Home;
