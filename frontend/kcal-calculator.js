var targetCalories = 0;
var totalKcal = 0;

        async function fetchFoods() {
            try {
                const response = await fetch('http://localhost:8080/foods');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const foods = await response.json();
                const foodList = document.getElementById('food-list');
    
            } catch (error) {
                console.error('Error fetching foods:', error);
            }
        }
            // Function to get target calorie
        async function getTargetCalories() {
            try {
                const response = await fetch('http://localhost:8080/targetkcal');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const targetKcal = await response.json();
                const targetElement = document.getElementById('target');
                targetElement.textContent = `Target Calories: ${targetKcal.targetKcal}`;
                targetCalories = targetKcal.targetKcal;
            } catch (error) {
                console.error('Error fetching target calories:', error);
            }
        }

        // Function to add a new food
        async function addFood(event) {
            event.preventDefault();
            const name = document.getElementById('name').value;
            const calories = parseInt(document.getElementById('calories').value, 10);
        
            try {
                const response = await fetch('http://localhost:8080/newfood', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ name, calories })
                });
        
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
        
                fetchFoods();
                document.getElementById('food-form').reset();
                totalKcal = targetCalories - calories;
                const totalElement = document.getElementById('total');
                totalElement.textContent = `You have eat today this much calories: ${totalKcal}`;
            } catch (error) {
                console.error('Error adding food:', error);
                alert('Failed to add food');
            }
    } 

        
        // Fetch foods when the page loads
        document.addEventListener('DOMContentLoaded', () => {
            getTargetCalories();
            fetchFoods()});
        
        // Add event listener to the form
        document.getElementById('food-form').addEventListener('submit', addFood);