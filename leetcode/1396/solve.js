
var UndergroundSystem = function () {
    this.checkInData = new Map();
    this.routeData = new Map();
};

UndergroundSystem.prototype.route = (start, end) => { return `${start}#${end}`; };

/** 
 * @param {number} id 
 * @param {string} stationName 
 * @param {number} t
 * @return {void}
 */
UndergroundSystem.prototype.checkIn = function (id, stationName, t) {
    this.checkInData.set(id, { checkInTime: t, checkInStation: stationName });
};

/** 
 * @param {number} id 
 * @param {string} stationName 
 * @param {number} t
 * @return {void}
 */
UndergroundSystem.prototype.checkOut = function (id, stationName, t) {
    const candidateCheckInData = this.checkInData.get(id);
    let routeProgress = this.routeData.get(this.route(candidateCheckInData.checkInStation, stationName));

    if (!routeProgress) {
        routeProgress = {
            time: 0,
            trips: 0,
        };
        this.routeData.set(this.route(candidateCheckInData.checkInStation, stationName), routeProgress);
    }

    routeProgress.trips++;
    routeProgress.time += (t - candidateCheckInData.checkInTime);
    this.checkInData.delete(id);
};

/** 
 * @param {string} startStation 
 * @param {string} endStation
 * @return {number}
 */
UndergroundSystem.prototype.getAverageTime = function (startStation, endStation) {
    const data = this.routeData.get(this.route(startStation, endStation));
    return data.time / data.trips;
};

/** 
 * Your UndergroundSystem object will be instantiated and called as such:
 * var obj = new UndergroundSystem()
 * obj.checkIn(id,stationName,t)
 * obj.checkOut(id,stationName,t)
 * var param_3 = obj.getAverageTime(startStation,endStation)
 */