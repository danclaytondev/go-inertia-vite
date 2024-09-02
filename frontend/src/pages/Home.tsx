import { Head } from "@inertiajs/react";
import Layout from "../components/Layout";

export default function Home() {
	return (
		<Layout>
			<div className="py-6">
				<Head title="Welcome" />
				<h1>Welcome</h1>
				<p>Hello, welcome to your first Inertia app!</p>
			</div>
		</Layout>
	);
}
