pragma solidity ^0.4.6;
contract VinTracker {
    
    uint index;
    function VinTracker() {
        index = 0;
    }

    struct VinEvent{
        uint previousIndex;
        uint index;
        string vin;
        address executor;
        string action;
    }
    
    mapping(uint => VinEvent) eventHistory;
    mapping(string => VinEvent) vinEvents;
    mapping(string => bool) vinExists;
    
    function addEvent(string vin, string action) returns (uint) {
        uint existingIndex;
        if (vinExists[vin]) {
            existingIndex = vinEvents[vin].index;
        } else {
            vinExists[vin] = true;
            existingIndex = 0;
        }
        
        vinEvents[vin] = VinEvent({previousIndex: existingIndex, index: ++index, vin: vin, executor: msg.sender,  action: action});
        eventHistory[index] = vinEvents[vin];
        return index;
    }
    
    function getEventById(uint history) returns (string) {
        PrintString(eventHistory[history].action);
        return eventHistory[history].action;
    }
    
    // TODO: Make it return arrays of event
    function getEventsByVIN(string vin) {
        if (vinExists[vin]) {
            string action = vinEvents[vin].action;
            PrintString(action);
            
            uint previousIndex = vinEvents[vin].previousIndex;
            
            while (previousIndex > 0) {
                PrintString(eventHistory[previousIndex].action);
                previousIndex = eventHistory[previousIndex].previousIndex;
            }
        } else {
            PrintString("not found");
        }
    }
    
    event PrintInt (uint number);
    event PrintString(string str);

    
    

    
}
