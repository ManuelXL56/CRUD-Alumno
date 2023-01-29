import { useEffect } from "react";

export default function Login() {

  // This hook is listening an event that came from the Iframe
  useEffect(() => {
    const handler = (ev: MessageEvent<{ type: string, message: string }>) => {
      //console.log('ev', ev)

      if (typeof ev.data !== 'object') return
      if (!ev.data.message) return

      //console.log('message: ', ev.data.message)
      localStorage.setItem("Session", ev.data.message);
      window.location.pathname = 'Dashboard';
    }

    window.addEventListener('message', handler)
    
    // Don't forget to remove addEventListener
    return () => window.removeEventListener('message', handler)
  }, [])

  return (
    <iframe title="Login" src="https://localhost/Login/" style={{width: "100%", height:"100%", border: "none"}}/>
  );
}