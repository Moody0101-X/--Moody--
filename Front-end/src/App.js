import "./css/index.css";
import "./css/global_fallback.css";
import { Routes, Route } from "react-router-dom";
import Home from "./Routes/Home";
import NavBar from "./Components/NavBar";
import Store from './Routes/Store';
import { login } from "./Store/UserStore";
import { SignUpRoute, LoginRoute } from "./Routes/AuthRoutes";
import { getJwtAuthToken, AuthUserWithJWT } from "./Server/Auth";
import { ApplicationMessage } from "./Components/UtilityComponents"
import { useEffect } from "react";
import { useSelector, useDispatch } from 'react-redux';
import { useState } from "react";


const App = () => {    
    const User = useSelector(state => state.User);
    const dispatch = useDispatch();
    const [Notification, setNotification] = useState(null);
    const [AppMessageId, SetAppMessageId] = useState(0);
    
    const dispatchAppMessageEvent = (msg, Case) => {
        
        setNotification({
            msg,
            Case,
            id: AppMessageId
       });

        SetAppMessageId(p => p + 1);
    }

    useEffect(() => {
                
        var RJWT = getJwtAuthToken();
        console.log(RJWT);
        
        if(RJWT) {
            AuthUserWithJWT()
            .then(response =>  response.json())
            .then(jsonData => {
                console.log(RJWT);
                if (jsonData.code === 200) {
                    const User = jsonData.data;
                    console.log(User);
                    dispatch(login(User));
                    return
                }
                
                dispatchAppMessageEvent(jsonData.data, "error");
            })
            .catch(e => console.log("ERR: ", e));
        }

    }, [])

    return (
        <>
            <NavBar  />
            <Routes>                
                <Route 
                    path="/Signup"
                    element={<SignUpRoute NotificationFunc={dispatchAppMessageEvent} />}
                />
                
                
                <Route 
                    path="/Login"
                    element={<LoginRoute NotificationFunc={dispatchAppMessageEvent} />}
                />
                
                <Route 
                    path="/" 
                    element={<Home NotificationFunc={dispatchAppMessageEvent} />} 
                />

                <Route 
                    path="/Home" 
                    element={<Home NotificationFunc={dispatchAppMessageEvent} />}
                />

                (User) ? (
                    
                    <Route 
                        path="/profile" 
                        element={<Home NotificationFunc={dispatchAppMessageEvent} />}
                    />

                ) : ()

                <Route
                    path="/Store"
                    element={<Store NotificationFunc={dispatchAppMessageEvent} />}
                />

    
           </ Routes>

           { (Notification) ? <ApplicationMessage msg={Notification.msg} Case={Notification.Case} id={Notification.id}/> : ("") }
        </>
    )
}

export default App;