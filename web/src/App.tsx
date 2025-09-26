import { Route, Router } from "@solidjs/router";
import { Layout } from "./components/root-layout";
import { Old } from "./Old";
import { Front } from "./Front";
import { Another } from "./Another";

export default function App() {
	return (
		<Router root={Layout}>
			<Route path="/" component={Front} />
			<Route path="/old" component={Old} />
			<Route path="/game" component={Another} />
		</Router>
	);
}
