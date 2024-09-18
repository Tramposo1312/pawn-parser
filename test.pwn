#include <a_samp>

#define MAX_PLAYERS 50

new gPlayerCount = 0;

public OnGameModeInit()
{
    print("Game mode initialized!");
    return 1;
}

public OnPlayerConnect(playerid)
{
    SendClientMessageToAll(0xFFFFFFAA, joinMessage);
    return 1;
}