"use client";

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { BACKEND_URL } from '@/loadEnv';

export default function Page() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [isAuthenticated, setIsAuthenticated] = useState(false); // Add isAuthenticated state

  useEffect(() => {
    const checkUserLoggedIn = async () => {
      const token = localStorage.getItem('authToken');
      if (!token) {
        router.push('/signin');
      } else {
        try {
          const response = await fetch(`${BACKEND_URL}/auth/verifyToken`, {
            headers: {
              'Authorization': `Bearer ${token}`
            }
          });

          if (!response.ok) {
            router.push('/auth/routes/signin');
          } else {
            setIsAuthenticated(true);
          }
        } catch (error) {
          console.error('Failed to verify token:', error);
          router.push('/signin');
        }
      }
      setIsLoading(false);
    };

    checkUserLoggedIn();
  }, [router]);

  if (isLoading) {
    return (
      <div style={{ backgroundColor: 'white', display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <div style={{ fontSize: '24px', fontWeight: 'bold' }}>
          Loading... <span role="img" aria-label="Loading">⌛️</span>
        </div>
      </div>
    );
  }
  /*
  if (!isAuthenticated) { // Add conditional rendering
    return null; // Return nothing before authentication check
  }
*/
//window.location.href = 'https://play.workadventure.localhost';

  return (
    <div> 
    </div>
  );
}
