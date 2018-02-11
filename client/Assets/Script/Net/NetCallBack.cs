using UnityEngine;
using System.Collections.Generic;

public class NetCallBack
{
    private static NetCallBack instance_ = null;

    public static NetCallBack Instance
    {
        get
        {
            if (instance_ == null)
                instance_ = new NetCallBack();
            return instance_;
        }
    }

    public void RegisterRsp(RegisterRsp rsp)
    {
        TestManager.Instance.RegisterRsp(rsp);
    }

    public void LoginRsp(LoginRsp rsp)
    {
        TestManager.Instance.LoginRsp(rsp);
    }

    public void FriendRsp(FriendRsp rsp)
    {
        Debug.Log("FriendRsp >>> " + rsp.Ok + " " + rsp.IdList.Count + " " + rsp.FriendList.Count);
    }

    public void RUOKRsp(RUOKRsp rsp)
    {

    }

    public void UserIntoMapRsp(UserIntoMapRsp rsp)
    {
        TestManager.Instance.UserIntoMap(rsp);
    }

    public void EntityIntoMapRsp(EntityIntoMapRsp rsp)
    {
        TestManager.Instance.EntityIntoMap(rsp);
    }

    public void EntityPosRsp(EntityPosRsp rsp)
    {
        TestManager.Instance.UpdateEntityPos(rsp);
    }

    public void EntityOutMapRsp(EntityOutMapRsp rsp)
    {
        TestManager.Instance.EntityOutMap(rsp);
    }

    public void EntityHpRsp(EntityHpRsp rsp)
    {
        TestManager.Instance.EntityHp(rsp);
    }

    public void BombInRsp(BombInRsp rsp)
    {
        TestManager.Instance.BombIn(rsp);
    }

    public void BombBlastRsp(BombBlastRsp rsp)
    {
        TestManager.Instance.BombBlast(rsp);
    }
}
