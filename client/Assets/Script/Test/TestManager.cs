using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class TestManager
{

    private static TestManager instance_ = null;

    private int userid_ = 0;
    public LoginObj currLoginObj_ = null;
    public List<LoginObj> lsLoginObj_ = new List<LoginObj>();
    public Dictionary<int, Entity> dictEntity_ = new Dictionary<int, Entity>();
    public Dictionary<int, Bomb> dictBomb_ = new Dictionary<int, Bomb>();

    public static TestManager Instance
    {
        get
        {
            if (instance_ == null)
                instance_ = new TestManager();
            return instance_;
        }
    }

    public void LoginRsp(LoginRsp rsp)
    {
        if (!rsp.OK)
        {
            if (currLoginObj_ != null)
            {
                NetSend.Instance.SendRegisterAsk(currLoginObj_.name_, currLoginObj_.passwd_);
            }
            else
            {
                Debug.Log("currLoginObj_ is null");
            }
        }
        else
        {
            userid_ = rsp.Userid;
            foreach (LoginObj obj in lsLoginObj_)
            {
                obj.Destroy();
            }
        }
    }

    public void RegisterRsp(RegisterRsp rsp)
    {
        if (rsp.OK)
        {
        }
    }

    public void UserIntoMap(UserIntoMapRsp rsp)
    {
        GameObject obj = GameObject.CreatePrimitive(PrimitiveType.Sphere);
        Entity entity = obj.AddComponent<Entity>();
        entity.Init(ENTITY_TYPE.USER, userid_, rsp.X, rsp.Y);

        obj.AddComponent<Control>();

        foreach (EntityInfo info in rsp.EntityList)
        {
            if (info.Entityid == userid_)
                continue;

            GameObject o;
            if (info.EntityType == ENTITY_TYPE.USER)
            {
                o = GameObject.CreatePrimitive(PrimitiveType.Cube);
            }
            else
            {
                o = GameObject.CreatePrimitive(PrimitiveType.Cylinder);
            }
            
            Entity e = o.AddComponent<Entity>();
            e.Init(info.EntityType, info.Entityid, info.X, info.Y);
            dictEntity_.Add(info.Entityid, e);
        }
    }

    public void EntityIntoMap(EntityIntoMapRsp rsp)
    {
        int entityid = rsp.Entity.Entityid;
        int x = rsp.Entity.X;
        int y = rsp.Entity.Y;
        int type = rsp.Entity.EntityType;
        if (entityid == userid_)
            return;

        if (!dictEntity_.ContainsKey(entityid))
        {
            GameObject o;
            if (type == ENTITY_TYPE.USER)
            {
                o = GameObject.CreatePrimitive(PrimitiveType.Cube);
            }
            else
            {
                o = GameObject.CreatePrimitive(PrimitiveType.Cylinder);
            }
            Entity e = o.AddComponent<Entity>();
            e.Init(type, entityid, x, y);
            dictEntity_.Add(entityid, e);
        }
    }

    public void EntityOutMap(EntityOutMapRsp rsp)
    {
        if (dictEntity_.ContainsKey(rsp.Entityid))
        {
            dictEntity_[rsp.Entityid].Destroy();
        }
    }

    public void UpdateEntityPos(EntityPosRsp rsp)
    {
        foreach (EntityInfo info in rsp.EntityList)
        {
            int entityid = info.Entityid;
            int x = info.X;
            int y = info.Y;
            int type = info.EntityType;
            if (entityid == userid_)
                return;

            if (!dictEntity_.ContainsKey(entityid))
            {
                GameObject o;
                if (type == ENTITY_TYPE.USER)
                {
                    o = GameObject.CreatePrimitive(PrimitiveType.Cube);
                }
                else
                {
                    o = GameObject.CreatePrimitive(PrimitiveType.Cylinder);
                }
                Entity e = o.AddComponent<Entity>();
                e.Init(type, entityid, x, y);
                dictEntity_.Add(entityid, e);
            }
            else
            {
                dictEntity_[entityid].UpdatePos(x, y);
            }
        }
    }

    public void EntityHp(EntityHpRsp rsp)
    {
        if (dictEntity_.ContainsKey(rsp.Entityid))
        {
            dictEntity_[rsp.Entityid].UpdateHp(rsp.Hp);
        }
    }

    public void BombIn(BombInRsp rsp)
    {
        if (!dictBomb_.ContainsKey(rsp.Bombid))
        {
            GameObject o = GameObject.CreatePrimitive(PrimitiveType.Sphere);
            Bomb b = o.AddComponent<Bomb>();
            b.Init(rsp.Bombid, rsp.X, rsp.Y);
            dictBomb_.Add(rsp.Bombid, b);
        }
    }

    public void BombBlast(BombBlastRsp rsp)
    {
        if (dictBomb_.ContainsKey(rsp.Bombid))
        {
            dictBomb_[rsp.Bombid].Blast();
            dictBomb_.Remove(rsp.Bombid);
        }
    }
}
