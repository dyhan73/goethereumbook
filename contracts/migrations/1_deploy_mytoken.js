var MyContract = artifacts.require("MyToken2");

module.exports = function(deployer) {
    // deployment steps
    deployer.deploy(MyContract);
};