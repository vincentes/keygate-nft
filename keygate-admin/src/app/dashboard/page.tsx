import { Button, Card } from "antd";

export default async function Dashboard() {
  return (
    <main className="flex min-h-screen flex-col items-left p-16 space-y-8">
      <div>
        <h1 className="text-4xl font-bold">Welcome!</h1>
        <p className="text-lg">What are you building today?</p>
      </div>
      <div className="flex flex-col items-left">
        <h2 className="text-2xl font-bold underline underline-offset-8 decoration-4 leading-loose">
          Use the core building blocks
        </h2>
        <div className="flex flex-row space-x-12">
          <Card title="Templates" className="w-96">
            <p>Design rewards and items for your customers.</p>
            <Button type="primary" className="mt-4">
              Create
            </Button>
          </Card>
          <Card title="Collectibles" className="w-96">
            <p>Use different designs to collectibles.</p>
            <Button type="primary" className="mt-4">
              Create
            </Button>
          </Card>
          <Card title="Release" className="w-96">
            <p>Release collectibles to your intended audience.</p>
            <Button type="primary" className="mt-4">
              Create
            </Button>
          </Card>
        </div>
      </div>
      <div className="flex flex-col items-left">
        <h2 className="text-2xl font-bold underline underline-offset-8 decoration-4 leading-loose">
          Create your own flows
        </h2>
        <div className="flex flex-row space-x-12">
          <Card title="Loyalty" className="w-96">
            <p>Connect Shopify and distribute rewards.</p>
            <Button>See example</Button>
          </Card>
          <Card title="Marketplace" className="w-96">
            <p>Create your own marketplace.</p>
            <Button>See example</Button>
          </Card>
          <Card title="Passport" className="w-96">
            <p>Enable features based on different rules.</p>
            <Button>See example</Button>
          </Card>
        </div>
      </div>
    </main>
  );
}
