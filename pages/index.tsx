import Link from 'next/link'
import Layout from '../components/Layout'
import { useState } from 'react'
import { CounterService } from '../sst-js/counter_connect'
import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
})

const client = createPromiseClient(CounterService, transport);

const IndexPage = () => {
  const [countTo, setCountTo] = useState(25);
  const [count, setCount] = useState(0);

  const submit = async () => {
    console.log("Submitting", countTo);
    
    const res = client.countTo({
      to: countTo
    })

    for await (const response of res) {
      setCount(response.count);
    }
  }

  return <Layout title="Home | Next.js + TypeScript Example">
    <input value={countTo} onChange={(e) => setCountTo(parseInt(e.target.value))} />
    <button onClick={submit}>Go!</button><br/>
  
    Current value: {count}
  </Layout>
}

export default IndexPage
