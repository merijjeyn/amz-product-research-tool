const ENDPOINT = 'http://localhost:8080'

async function login(email, name, sub) {
    const res = await fetch(ENDPOINT + '/login', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin':'*',
        },
        body: JSON.stringify({ 
            'email': email,
            'name': name, 
            'gid': sub
        })
    })
    console.log(res.status)

    return res.ok;
}

export { login };
