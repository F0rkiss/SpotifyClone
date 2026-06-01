
import { createBrowserRouter } from 'react-router-dom'
import ProtectedRoute from './ProtectedRoute';
// IMPORT PAGE YANG ADA 
import LoginPage from '../pages/LoginPage';
import TestPage from '../pages/TestPage';

const router = createBrowserRouter([
    {path:"/login", element:<LoginPage/>},
    {
        path:"/", 
        element:(
            <ProtectedRoute>
               <TestPage /> 
            </ProtectedRoute>
        )
    },





// Next idea when User DIDN'T LOGIN! but still want to play music
    // {
    //     path:"/", 
    //     element:(
    //         <TestPage /> 
    //     )
    // },
]);

export default router