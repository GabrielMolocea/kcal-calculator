import { useEffect, useState } from "react";
import CaloiresForm from "./components/form/caloriesForm";

function App() {
	const [targetCalories, setTargetCalories] = useState(0);
	const [totalKcal, setTotalKcal] = useState(0);
	const [foodList, setFoodList] = useState<
		{ id: number; name: string; calories: number }[]
	>([]);

	useEffect(() => {
		getTargetCalories();
		getFoodList();
	}, []);

	async function getTargetCalories() {
		await fetch("http://localhost:8080/targetkcal")
			.then((response) => response.json())
			.then((data) => setTargetCalories(data.targetKcal));
	}

	async function getFoodList() {
		const response = await fetch("http://localhost:8080/foods", {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
			},
		});
		setFoodList(await response.json());
	}

	console.log(foodList);

	async function resetTargetCaloris() {
		await fetch("http://localhost:8080/resttargetkcal", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
		})
			.then((res) => res.json())
			.then((data) => setTargetCalories(data.targetKcal))
			.then(() => getTargetCalories());
	}

	return (
		<>
			<div id="app">
				<h1>Calorie Calculator</h1>
				<h2 id="target">Target Calories: {targetCalories} </h2>
				<CaloiresForm setTotalKcal={setTotalKcal} foodList={foodList} />
				<h2 id="total">You have eaten this much calories today: {totalKcal}</h2>

				<button onClick={resetTargetCaloris}>Reset target</button>
			</div>
		</>
	);
}

export default App;
