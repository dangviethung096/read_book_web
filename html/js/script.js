function getDayLeft() {
    var now = new Date();
    var tetHoliday = new Date(2024, 1, 14);
    var differenceInMilliseconds = tetHoliday - now;
    var differenceInDays = Math.ceil(differenceInMilliseconds / (1000 * 60 * 60 * 24));

    document.getElementById("dayLeft").innerHTML = `${differenceInDays} days left`;
}